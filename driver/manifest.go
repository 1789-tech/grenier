package driver

import (
	"fmt"
	"strconv"
	"strings"
)

// Manifest is the machine-readable, review-surface description a driver ships
// next to its code. It carries the driver's permission posture, detection risk,
// run cap, and — critically — the accessible-name matchers for the draft and
// publish actions, each with a confidence note and a last-verified date. The
// manifest is part of the review surface: if a driver is ToS-hostile it says so
// here, so the core stays neutral and each driver carries its own posture.
type Manifest struct {
	ID                string
	Country           string
	DisplayName       string
	Surface           string // "full_browser", "official_api"
	PermissionPosture PermissionPosture
	DetectionRisk     RiskLevel
	MaxItemsPerRun    int
	LastVerified      string

	DraftSupported     bool
	TerminalAction     string // e.g. "save_draft"
	PublishActionWired bool   // must be false in V1

	NewListingURL string

	// Accessible-name matchers used to disambiguate the terminal draft action
	// from publish. Lower-cased and trimmed on load.
	DraftActionNames   []string
	PublishActionNames []string
	SelectorConfidence string // "platform-confirmed" | "vendor-confirmed" | "field-report"

	Sources map[string]string
}

// Validate enforces the V1 product contract at load time. A manifest that would
// let a driver publish, or that has no draft stop, is rejected before any
// browser work.
func (m Manifest) Validate() error {
	if m.ID == "" {
		return fmt.Errorf("manifest: id is required")
	}
	if m.PublishActionWired {
		return fmt.Errorf("manifest %s: publish_action_wired must be false in V1 (save-as-draft is terminal)", m.ID)
	}
	if !m.DraftSupported || m.TerminalAction == "" {
		return fmt.Errorf("manifest %s: no stable draft stop — not eligible for a V1 driver (research note only)", m.ID)
	}
	if len(m.DraftActionNames) == 0 {
		return fmt.Errorf("manifest %s: no draft action names declared — cannot identify the terminal control", m.ID)
	}
	if m.LastVerified == "" {
		return fmt.Errorf("manifest %s: last_verified date is required (selectors drift)", m.ID)
	}
	switch m.SelectorConfidence {
	case "platform-confirmed", "vendor-confirmed", "field-report":
	case "":
		return fmt.Errorf("manifest %s: selector confidence is required", m.ID)
	default:
		return fmt.Errorf("manifest %s: unknown selector confidence %q", m.ID, m.SelectorConfidence)
	}
	return nil
}

// ParseManifest loads a driver manifest from its TOML bytes. It uses a tiny
// dependency-free TOML subset reader — enough for the flat keys and the [draft],
// [selectors], and [sources] tables the manifest schema uses — so the core has
// zero external dependencies.
func ParseManifest(data []byte) (Manifest, error) {
	tables, err := parseTOMLSubset(data)
	if err != nil {
		return Manifest{}, err
	}
	root := tables[""]
	draft := tables["draft"]
	sel := tables["selectors"]

	m := Manifest{
		ID:                 root.str("id"),
		Country:            root.str("country"),
		DisplayName:        root.str("display_name"),
		Surface:            root.str("surface"),
		PermissionPosture:  PermissionPosture(root.str("permission_posture")),
		DetectionRisk:      RiskLevel(root.str("detection_risk")),
		MaxItemsPerRun:     root.int("max_items_per_run"),
		LastVerified:       root.str("last_verified"),
		DraftSupported:     draft.bool("supported"),
		TerminalAction:     draft.str("terminal_action"),
		PublishActionWired: draft.bool("publish_action_wired"),
		NewListingURL:      sel.str("new_listing_url"),
		DraftActionNames:   lowerAll(sel.strs("draft_action_names")),
		PublishActionNames: lowerAll(sel.strs("publish_action_names")),
		SelectorConfidence: sel.str("confidence"),
		Sources:            tables["sources"].asStringMap(),
	}
	if m.PermissionPosture == "" {
		m.PermissionPosture = PermissionUnknown
	}
	if m.DetectionRisk == "" {
		m.DetectionRisk = RiskUnknown
	}
	return m, nil
}

func lowerAll(in []string) []string {
	out := make([]string, len(in))
	for i, s := range in {
		out[i] = strings.ToLower(strings.TrimSpace(s))
	}
	return out
}

// --- minimal TOML subset reader ---------------------------------------------
//
// Supports: # comments, [table] headers, and key = value where value is a
// double-quoted string, a bool, an integer, or a single-line array of quoted
// strings. This is all the manifest schema uses; anything richer belongs in a
// real TOML library, which we intentionally avoid to keep the core dep-free.

type tomlTable map[string]any

func (t tomlTable) str(k string) string {
	if v, ok := t[k].(string); ok {
		return v
	}
	return ""
}

func (t tomlTable) bool(k string) bool {
	b, _ := t[k].(bool)
	return b
}

func (t tomlTable) int(k string) int {
	if n, ok := t[k].(int64); ok {
		return int(n)
	}
	return 0
}

func (t tomlTable) strs(k string) []string {
	if v, ok := t[k].([]string); ok {
		return v
	}
	return nil
}

func (t tomlTable) asStringMap() map[string]string {
	out := map[string]string{}
	for k, v := range t {
		if s, ok := v.(string); ok {
			out[k] = s
		}
	}
	return out
}

func parseTOMLSubset(data []byte) (map[string]tomlTable, error) {
	tables := map[string]tomlTable{"": {}}
	current := ""
	for lineNo, raw := range strings.Split(string(data), "\n") {
		line := stripComment(raw)
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			current = strings.TrimSpace(line[1 : len(line)-1])
			if _, ok := tables[current]; !ok {
				tables[current] = tomlTable{}
			}
			continue
		}
		eq := strings.IndexByte(line, '=')
		if eq < 0 {
			return nil, fmt.Errorf("toml line %d: no '=' in %q", lineNo+1, line)
		}
		key := strings.TrimSpace(line[:eq])
		val := strings.TrimSpace(line[eq+1:])
		parsed, err := parseTOMLValue(val)
		if err != nil {
			return nil, fmt.Errorf("toml line %d (%s): %w", lineNo+1, key, err)
		}
		tables[current][key] = parsed
	}
	return tables, nil
}

// stripComment removes a trailing # comment that is not inside a quoted string.
func stripComment(s string) string {
	inStr := false
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '"':
			inStr = !inStr
		case '#':
			if !inStr {
				return s[:i]
			}
		}
	}
	return s
}

func parseTOMLValue(v string) (any, error) {
	switch {
	case v == "true":
		return true, nil
	case v == "false":
		return false, nil
	case strings.HasPrefix(v, "\""):
		return unquoteTOML(v)
	case strings.HasPrefix(v, "["):
		return parseTOMLArray(v)
	default:
		n, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("unsupported value %q", v)
		}
		return n, nil
	}
}

func unquoteTOML(v string) (string, error) {
	if len(v) < 2 || !strings.HasPrefix(v, "\"") || !strings.HasSuffix(v, "\"") {
		return "", fmt.Errorf("unterminated string %q", v)
	}
	return v[1 : len(v)-1], nil
}

func parseTOMLArray(v string) ([]string, error) {
	if !strings.HasPrefix(v, "[") || !strings.HasSuffix(v, "]") {
		return nil, fmt.Errorf("unterminated array %q", v)
	}
	inner := strings.TrimSpace(v[1 : len(v)-1])
	if inner == "" {
		return []string{}, nil
	}
	var out []string
	for _, part := range splitTopComma(inner) {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		s, err := unquoteTOML(part)
		if err != nil {
			return nil, err
		}
		out = append(out, s)
	}
	return out, nil
}

// splitTopComma splits on commas that are not inside quotes.
func splitTopComma(s string) []string {
	var parts []string
	var cur strings.Builder
	inStr := false
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '"':
			inStr = !inStr
			cur.WriteByte(s[i])
		case ',':
			if inStr {
				cur.WriteByte(s[i])
			} else {
				parts = append(parts, cur.String())
				cur.Reset()
			}
		default:
			cur.WriteByte(s[i])
		}
	}
	if cur.Len() > 0 {
		parts = append(parts, cur.String())
	}
	return parts
}
