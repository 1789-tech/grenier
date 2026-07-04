package main

import (
	"os"
	"path/filepath"
	"testing"
)

// loadSuite must accept the real fixtures and reject duplicate ids.
func TestLoadRealSuite(t *testing.T) {
	suite, err := loadSuite("../../evals/evals.json")
	if err != nil {
		t.Fatal(err)
	}
	if len(suite.Cases) == 0 {
		t.Fatal("no cases loaded")
	}
	for _, tc := range suite.Cases {
		if tc.Skill != "grenier" {
			t.Errorf("case %q targets skill %q; expected the single-skill layout (grenier)", tc.ID, tc.Skill)
		}
	}
}

func TestLoadSuiteRejectsDuplicateIDs(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "evals.json")
	os.WriteFile(path, []byte(`{"cases":[
		{"id":"a","skill":"grenier","prompt":"x","must_include":["y"]},
		{"id":"a","skill":"grenier","prompt":"x","must_include":["y"]}
	]}`), 0o644)
	if _, err := loadSuite(path); err == nil {
		t.Fatal("expected duplicate id to fail")
	}
}

func TestLoadSuiteRequiresAssertion(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "evals.json")
	os.WriteFile(path, []byte(`{"cases":[
		{"id":"a","skill":"grenier","prompt":"x","must_include":[],"must_avoid":[]}
	]}`), 0o644)
	if _, err := loadSuite(path); err == nil {
		t.Fatal("expected case with no assertions to fail")
	}
}

// skillContext must bundle the core SKILL.md plus the progressive-disclosure
// files (references/ + countries/) so with-skill scoring sees the full skill.
func TestSkillContextBundlesProgressiveDisclosure(t *testing.T) {
	ctx, err := skillContext("../../skills", "grenier")
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]bool{
		"references/pricing.md":     false,
		"references/listing.md":     false,
		"references/disposition.md": false,
		"countries/fr.md":           false,
		"countries/uk.md":           false,
		"countries/us.md":           false,
	}
	for _, e := range ctx.extras {
		if _, ok := want[e]; ok {
			want[e] = true
		}
	}
	for f, seen := range want {
		if !seen {
			t.Errorf("expected %s in bundle, got extras=%v", f, ctx.extras)
		}
	}
	for _, needle := range []string{"# skill: grenier", "# references/pricing.md", "# countries/fr.md"} {
		if !contains(ctx.bundle, needle) {
			t.Errorf("bundle missing header %q", needle)
		}
	}
}

func TestSkillContextMissingSkill(t *testing.T) {
	if _, err := skillContext("../../skills", "does-not-exist"); err == nil {
		t.Fatal("expected missing skill to error")
	}
}

func TestScore(t *testing.T) {
	tc := Case{
		MustInclude: []string{"Leboncoin", "35"},
		MustAvoid:   []string{"eBay", "Vinted"},
	}
	got := score(tc, "Poste sur Leboncoin à 35 euros, remise en main propre.")
	if got.Total != 4 {
		t.Fatalf("total = %d, want 4", got.Total)
	}
	if got.Passed != 4 { // both includes present, both avoids absent
		t.Fatalf("passed = %d, want 4", got.Passed)
	}

	bad := score(tc, "Mets-le sur eBay et Vinted.")
	if bad.Passed != 0 { // no includes, both avoids present
		t.Fatalf("passed = %d, want 0", bad.Passed)
	}
}

func contains(haystack, needle string) bool {
	return len(haystack) >= len(needle) && indexOf(haystack, needle) >= 0
}

func indexOf(s, sub string) int {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return i
		}
	}
	return -1
}
