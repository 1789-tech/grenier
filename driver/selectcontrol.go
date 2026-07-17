package driver

import "strings"

// TerminalChoice is the outcome of resolving which control, if any, is the safe
// terminal draft action on the page.
type TerminalChoice struct {
	Control Control     // the draft control to click (valid only when Status == "")
	Status  DraftStatus // non-empty means: do NOT click, skip with this status
	Reason  string
	Seen    []string // accessible names seen, for the audit log
}

// SelectTerminalControl implements the V1 safety rule for choosing the terminal
// action from the page's visible controls, using the manifest's declared draft
// and publish accessible-name matchers. It NEVER returns a publish control.
//
// Rules (contract §Safety model):
//   - draft control must match a marketplace-specific draft action name;
//   - if the draft control is not uniquely identified while a publish control is
//     visible, skip with skipped_ambiguous_submit;
//   - a page with a publish action but no draft action is publish-only:
//     disqualified (skipped_publish_only);
//   - no draft and no publish match means selector drift: skip, never
//     best-effort publish.
func SelectTerminalControl(controls []Control, m Manifest) TerminalChoice {
	var (
		draftHits   []Control
		publishSeen bool
		seen        []string
	)
	for _, c := range controls {
		name := strings.ToLower(strings.TrimSpace(c.AccessibleName))
		seen = append(seen, c.AccessibleName)
		isDraft := matchesAny(name, m.DraftActionNames)
		isPublish := matchesAny(name, m.PublishActionNames)
		if isPublish {
			publishSeen = true
		}
		// A control that matches BOTH a draft and a publish pattern is ambiguous
		// by construction — never treat it as a clean draft hit.
		if isDraft && !isPublish {
			draftHits = append(draftHits, c)
		}
	}

	switch {
	case len(draftHits) == 1:
		return TerminalChoice{Control: draftHits[0], Seen: seen}
	case len(draftHits) > 1:
		return TerminalChoice{
			Status: StatusSkippedAmbiguous,
			Reason: "multiple controls match the draft action name — cannot uniquely identify the draft stop",
			Seen:   seen,
		}
	case publishSeen:
		// draft hits == 0 but publish is present.
		return TerminalChoice{
			Status: StatusSkippedPublish,
			Reason: "publish control present but no draft stop found — publish-only page, disqualified",
			Seen:   seen,
		}
	default:
		return TerminalChoice{
			Status: StatusSkippedDrift,
			Reason: "no draft and no publish control matched the manifest — selector drift, skipping instead of guessing",
			Seen:   seen,
		}
	}
}

// matchesAny reports whether the (already lower-cased) accessible name matches
// any of the declared patterns. A pattern matches if the name equals it or
// contains it as a substring — declared patterns are specific labels, so this
// stays tight while tolerating surrounding whitespace/decoration.
func matchesAny(name string, patterns []string) bool {
	for _, p := range patterns {
		if p == "" {
			continue
		}
		if name == p || strings.Contains(name, p) {
			return true
		}
	}
	return false
}
