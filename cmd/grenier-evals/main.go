package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Suite struct {
	Cases []Case `json:"cases"`
}

type Case struct {
	ID         string   `json:"id"`
	Skill      string   `json:"skill"`
	Prompt     string   `json:"prompt"`
	MustInclude []string `json:"must_include"`
	MustAvoid   []string `json:"must_avoid"`
}

type Result struct {
	CaseID string
	Skill  string
	With   Score
	Without Score
}

type Score struct {
	Passed int
	Total  int
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	var evalPath, skillsDir, genCmd string
	flag.StringVar(&evalPath, "evals", "evals/evals.json", "eval fixture path")
	flag.StringVar(&skillsDir, "skills", "skills", "skills directory")
	flag.StringVar(&genCmd, "generate-cmd", os.Getenv("EVAL_GENERATE_CMD"), "optional shell command that reads a prompt on stdin and writes the model answer")
	flag.Parse()

	suite, err := loadSuite(evalPath)
	if err != nil {
		return err
	}

	if genCmd == "" {
		return validateSuite(suite, skillsDir)
	}

	var results []Result
	for _, tc := range suite.Cases {
		skillText, err := os.ReadFile(filepath.Join(skillsDir, tc.Skill, "SKILL.md"))
		if err != nil {
			return fmt.Errorf("%s: read skill: %w", tc.ID, err)
		}

		withOutput, err := generate(genCmd, string(skillText)+"\n\nUser:\n"+tc.Prompt)
		if err != nil {
			return fmt.Errorf("%s: with skill: %w", tc.ID, err)
		}
		withoutOutput, err := generate(genCmd, tc.Prompt)
		if err != nil {
			return fmt.Errorf("%s: without skill: %w", tc.ID, err)
		}

		results = append(results, Result{
			CaseID:  tc.ID,
			Skill:   tc.Skill,
			With:    score(tc, withOutput),
			Without: score(tc, withoutOutput),
		})
	}

	printSummary(results)
	return nil
}

func loadSuite(path string) (Suite, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Suite{}, err
	}
	var suite Suite
	if err := json.Unmarshal(data, &suite); err != nil {
		return Suite{}, err
	}
	if len(suite.Cases) == 0 {
		return Suite{}, fmt.Errorf("no eval cases in %s", path)
	}
	for _, tc := range suite.Cases {
		if tc.ID == "" || tc.Skill == "" || strings.TrimSpace(tc.Prompt) == "" {
			return Suite{}, fmt.Errorf("invalid eval case: id, skill, and prompt are required")
		}
		if len(tc.MustInclude)+len(tc.MustAvoid) == 0 {
			return Suite{}, fmt.Errorf("%s: at least one assertion is required", tc.ID)
		}
	}
	return suite, nil
}

func validateSuite(suite Suite, skillsDir string) error {
	bySkill := map[string]int{}
	for _, tc := range suite.Cases {
		if _, err := os.Stat(filepath.Join(skillsDir, tc.Skill, "SKILL.md")); err != nil {
			return fmt.Errorf("%s: read skill: %w", tc.ID, err)
		}
		bySkill[tc.Skill]++
	}

	fmt.Println("grenier evals")
	fmt.Printf("validated %d cases across %d skills\n", len(suite.Cases), len(bySkill))
	for skill, count := range bySkill {
		fmt.Printf("- %s: %d cases\n", skill, count)
	}
	fmt.Println("set EVAL_GENERATE_CMD or -generate-cmd to run with-skill / without-skill scoring.")
	return nil
}

func generate(command, prompt string) (string, error) {
	if command == "" {
		return "", nil
	}
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdin = strings.NewReader(prompt)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("%w: %s", err, strings.TrimSpace(stderr.String()))
	}
	return stdout.String(), nil
}

func score(tc Case, output string) Score {
	lower := strings.ToLower(output)
	passed := 0
	total := len(tc.MustInclude) + len(tc.MustAvoid)
	for _, needle := range tc.MustInclude {
		if strings.Contains(lower, strings.ToLower(needle)) {
			passed++
		}
	}
	for _, needle := range tc.MustAvoid {
		if !strings.Contains(lower, strings.ToLower(needle)) {
			passed++
		}
	}
	return Score{Passed: passed, Total: total}
}

func printSummary(results []Result) {
	var withPassed, withTotal, withoutPassed, withoutTotal int
	bySkill := map[string][4]int{}
	for _, res := range results {
		withPassed += res.With.Passed
		withTotal += res.With.Total
		withoutPassed += res.Without.Passed
		withoutTotal += res.Without.Total

		cur := bySkill[res.Skill]
		cur[0] += res.With.Passed
		cur[1] += res.With.Total
		cur[2] += res.Without.Passed
		cur[3] += res.Without.Total
		bySkill[res.Skill] = cur
	}

	fmt.Println("grenier evals")
	fmt.Printf("overall: with skill %.0f%% | without skill %.0f%% | delta %+.0f pts\n",
		percent(withPassed, withTotal), percent(withoutPassed, withoutTotal), percent(withPassed, withTotal)-percent(withoutPassed, withoutTotal))
	for skill, totals := range bySkill {
		fmt.Printf("- %s: with %.0f%% | without %.0f%% | delta %+.0f pts\n",
			skill, percent(totals[0], totals[1]), percent(totals[2], totals[3]), percent(totals[0], totals[1])-percent(totals[2], totals[3]))
	}
}

func percent(n, d int) float64 {
	if d == 0 {
		return 0
	}
	return float64(n) * 100 / float64(d)
}
