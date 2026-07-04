// Command grenier-evals is a lightweight eval runner for the grenier skill.
//
// grenier is a single, prompt-only skill with progressive disclosure: a core
// SKILL.md plus on-demand references/ and countries/ recipes. The fixtures in
// evals/evals.json are plain string assertions describing what a good answer
// must include or avoid.
//
// Two modes:
//
//   - validate (default): parse the fixtures, check every referenced skill and
//     its progressive-disclosure bundle exist, and report shape. No model call,
//     so it runs in CI with zero dependencies.
//   - score (-generate-cmd / EVAL_GENERATE_CMD set): for each case, generate an
//     answer WITH the full skill context and WITHOUT it, score both against the
//     must_include / must_avoid assertions, and print the with-vs-without delta.
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

type Suite struct {
	Cases []Case `json:"cases"`
}

type Case struct {
	ID          string   `json:"id"`
	Skill       string   `json:"skill"`
	Prompt      string   `json:"prompt"`
	Delta       string   `json:"delta"`
	MustInclude []string `json:"must_include"`
	MustAvoid   []string `json:"must_avoid"`
}

type Result struct {
	CaseID  string
	Skill   string
	With    Score
	Without Score
}

type Score struct {
	Passed int
	Total  int
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "grenier-evals:", err)
		os.Exit(1)
	}
}

func run() error {
	var evalPath, skillsDir, genCmd string
	flag.StringVar(&evalPath, "evals", "evals/evals.json", "eval fixture path")
	flag.StringVar(&skillsDir, "skills", "skills", "skills directory")
	flag.StringVar(&genCmd, "generate-cmd", os.Getenv("EVAL_GENERATE_CMD"), "shell command that reads a prompt on stdin and writes the model answer; enables with/without scoring")
	flag.Parse()

	suite, err := loadSuite(evalPath)
	if err != nil {
		return err
	}

	if genCmd == "" {
		return validateSuite(suite, skillsDir)
	}
	return scoreSuite(suite, skillsDir, genCmd)
}

func loadSuite(path string) (Suite, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Suite{}, err
	}
	var suite Suite
	if err := json.Unmarshal(data, &suite); err != nil {
		return Suite{}, fmt.Errorf("parse %s: %w", path, err)
	}
	if len(suite.Cases) == 0 {
		return Suite{}, fmt.Errorf("no eval cases in %s", path)
	}
	seen := map[string]bool{}
	for _, tc := range suite.Cases {
		if tc.ID == "" || tc.Skill == "" || strings.TrimSpace(tc.Prompt) == "" {
			return Suite{}, fmt.Errorf("invalid eval case: id, skill and prompt are required")
		}
		if seen[tc.ID] {
			return Suite{}, fmt.Errorf("duplicate eval id %q", tc.ID)
		}
		seen[tc.ID] = true
		if len(tc.MustInclude)+len(tc.MustAvoid) == 0 {
			return Suite{}, fmt.Errorf("%s: at least one must_include/must_avoid assertion is required", tc.ID)
		}
	}
	return suite, nil
}

// validateSuite checks structure only — no model call. It confirms every skill a
// case references exists and reports how many progressive-disclosure files
// (references/ + countries/) are bundled with it, so a fixture pointing at the
// old multi-skill taxonomy fails loudly against the current single-skill layout.
func validateSuite(suite Suite, skillsDir string) error {
	bySkill := map[string]int{}
	for _, tc := range suite.Cases {
		if _, err := skillContext(skillsDir, tc.Skill); err != nil {
			return fmt.Errorf("%s: %w", tc.ID, err)
		}
		bySkill[tc.Skill]++
	}

	fmt.Println("grenier evals — validate")
	fmt.Printf("validated %d cases across %d skill(s)\n", len(suite.Cases), len(bySkill))
	for _, skill := range sortedKeys(bySkill) {
		ctx, _ := skillContext(skillsDir, skill)
		fmt.Printf("- %s: %d case(s), %d supporting file(s) bundled\n", skill, bySkill[skill], len(ctx.extras))
	}
	fmt.Println("set EVAL_GENERATE_CMD (or -generate-cmd) to run with-skill / without-skill scoring.")
	return nil
}

func scoreSuite(suite Suite, skillsDir, genCmd string) error {
	var results []Result
	for _, tc := range suite.Cases {
		ctx, err := skillContext(skillsDir, tc.Skill)
		if err != nil {
			return fmt.Errorf("%s: %w", tc.ID, err)
		}
		withOut, err := generate(genCmd, ctx.bundle+"\n\nUser:\n"+tc.Prompt)
		if err != nil {
			return fmt.Errorf("%s: with skill: %w", tc.ID, err)
		}
		withoutOut, err := generate(genCmd, tc.Prompt)
		if err != nil {
			return fmt.Errorf("%s: without skill: %w", tc.ID, err)
		}
		results = append(results, Result{
			CaseID:  tc.ID,
			Skill:   tc.Skill,
			With:    score(tc, withOut),
			Without: score(tc, withoutOut),
		})
	}
	printSummary(results)
	return nil
}

type skillBundle struct {
	bundle string   // SKILL.md + all supporting markdown, concatenated
	extras []string // relative paths of the progressive-disclosure files
}

// skillContext assembles the full progressive-disclosure context a model would
// have access to: the core SKILL.md plus every other markdown file under the
// skill (references/, countries/). Failing to find SKILL.md is a hard error.
func skillContext(skillsDir, skill string) (skillBundle, error) {
	root := filepath.Join(skillsDir, skill)
	skillPath := filepath.Join(root, "SKILL.md")
	head, err := os.ReadFile(skillPath)
	if err != nil {
		return skillBundle{}, fmt.Errorf("read skill %q: %w", skill, err)
	}

	var extras []string
	var b strings.Builder
	b.WriteString("# skill: " + skill + "\n")
	b.Write(head)

	err = filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || path == skillPath || !strings.HasSuffix(path, ".md") {
			return nil
		}
		rel, _ := filepath.Rel(root, path)
		extras = append(extras, filepath.ToSlash(rel))
		return nil
	})
	if err != nil {
		return skillBundle{}, err
	}
	sort.Strings(extras)
	for _, rel := range extras {
		data, err := os.ReadFile(filepath.Join(root, rel))
		if err != nil {
			return skillBundle{}, err
		}
		b.WriteString("\n\n# " + rel + "\n")
		b.Write(data)
	}
	return skillBundle{bundle: b.String(), extras: extras}, nil
}

func generate(command, prompt string) (string, error) {
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

// score counts satisfied assertions: each must_include substring present and
// each must_avoid substring absent scores one point. Matching is case-folded.
func score(tc Case, output string) Score {
	lower := strings.ToLower(output)
	passed := 0
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
	return Score{Passed: passed, Total: len(tc.MustInclude) + len(tc.MustAvoid)}
}

func printSummary(results []Result) {
	var withP, withT, woutP, woutT int
	bySkill := map[string][4]int{}
	for _, res := range results {
		withP += res.With.Passed
		withT += res.With.Total
		woutP += res.Without.Passed
		woutT += res.Without.Total
		cur := bySkill[res.Skill]
		cur[0] += res.With.Passed
		cur[1] += res.With.Total
		cur[2] += res.Without.Passed
		cur[3] += res.Without.Total
		bySkill[res.Skill] = cur
	}

	fmt.Println("grenier evals — score")
	fmt.Printf("overall: with skill %.0f%% | without skill %.0f%% | delta %+.0f pts\n",
		percent(withP, withT), percent(woutP, woutT), percent(withP, withT)-percent(woutP, woutT))
	for _, skill := range sortedKeys4(bySkill) {
		t := bySkill[skill]
		fmt.Printf("- %s: with %.0f%% | without %.0f%% | delta %+.0f pts\n",
			skill, percent(t[0], t[1]), percent(t[2], t[3]), percent(t[0], t[1])-percent(t[2], t[3]))
	}
}

func percent(n, d int) float64 {
	if d == 0 {
		return 0
	}
	return float64(n) * 100 / float64(d)
}

func sortedKeys(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func sortedKeys4(m map[string][4]int) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
