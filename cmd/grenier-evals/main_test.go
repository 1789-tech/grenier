package main

import "testing"

func TestLoadSuite(t *testing.T) {
	suite, err := loadSuite("../../evals/evals.json")
	if err != nil {
		t.Fatal(err)
	}
	seen := map[string]bool{}
	for _, tc := range suite.Cases {
		if seen[tc.ID] {
			t.Fatalf("duplicate eval id %q", tc.ID)
		}
		seen[tc.ID] = true
	}
}
