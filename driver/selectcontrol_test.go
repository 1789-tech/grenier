package driver

import (
	"strings"
	"testing"
)

func testManifest() Manifest {
	return Manifest{
		ID:                 "test-fr",
		DraftActionNames:   []string{"enregistrer le brouillon", "sauvegarder le brouillon"},
		PublishActionNames: []string{"publier", "mettre en ligne"},
		SelectorConfidence: "field-report",
		LastVerified:       "2026-07-17",
		DraftSupported:     true,
		TerminalAction:     "save_draft",
	}
}

func TestSelectTerminalControl(t *testing.T) {
	m := testManifest()
	draft := Control{AccessibleName: "Enregistrer le brouillon", ID: "d"}
	draft2 := Control{AccessibleName: "Sauvegarder le brouillon", ID: "d2"}
	publish := Control{AccessibleName: "Publier", ID: "p"}

	tests := []struct {
		name     string
		controls []Control
		want     DraftStatus // "" == picked cleanly
		wantID   string
	}{
		{"unique draft with publish present", []Control{publish, draft}, "", "d"},
		{"unique draft alone", []Control{draft}, "", "d"},
		{"two drafts -> ambiguous", []Control{publish, draft, draft2}, StatusSkippedAmbiguous, ""},
		{"publish only -> disqualified", []Control{publish}, StatusSkippedPublish, ""},
		{"nothing matches -> drift", []Control{{AccessibleName: "Continuer"}}, StatusSkippedDrift, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SelectTerminalControl(tt.controls, m)
			if got.Status != tt.want {
				t.Fatalf("status = %q, want %q", got.Status, tt.want)
			}
			if tt.want == "" && got.Control.ID != tt.wantID {
				t.Fatalf("picked %q, want %q", got.Control.ID, tt.wantID)
			}
			// Invariant: a cleanly picked control is never a publish control.
			if tt.want == "" && matchesAny(strings.ToLower(got.Control.AccessibleName), m.PublishActionNames) {
				t.Fatalf("picked a publish control: %q", got.Control.AccessibleName)
			}
		})
	}
}

// A control that matches both a draft and a publish pattern must never be picked.
func TestSelectTerminalControl_OverlapNeverPicked(t *testing.T) {
	m := testManifest()
	m.DraftActionNames = append(m.DraftActionNames, "publier le brouillon") // pathological overlap
	overlap := Control{AccessibleName: "Publier le brouillon", ID: "x"}
	got := SelectTerminalControl([]Control{overlap}, m)
	if got.Status == "" {
		t.Fatalf("overlapping draft/publish control was picked: %+v", got.Control)
	}
}
