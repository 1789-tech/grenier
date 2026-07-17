package driver

import "context"

// Page is the minimal browser port a driver needs. A real implementation drives
// a user-owned Chrome/Firefox profile (chromedp, rod, or a browser extension via
// native messaging); the fixturemarket package implements it against an
// in-memory fake DOM so drivers can be proven end-to-end without touching a real
// account. The interface is identical for both, so the same driver code runs
// against a fixture in CI and a real browser in production.
type Page interface {
	// Navigate loads a URL (the driver's new-listing page).
	Navigate(ctx context.Context, url string) error
	// UploadPhotos attaches local files to the control matched by field.
	UploadPhotos(ctx context.Context, field FieldRef, paths []string) error
	// Fill types value into the control matched by field.
	Fill(ctx context.Context, field FieldRef, value string) error
	// Controls returns the currently visible, actionable controls (buttons,
	// links, submits) with their accessible names. This is how a driver
	// discovers — and disambiguates — the draft vs publish actions.
	Controls(ctx context.Context) ([]Control, error)
	// Click activates one previously discovered control.
	Click(ctx context.Context, c Control) error
	// Snapshot captures the accessibility tree / a screenshot handle just before
	// a terminal action, for the audit log. Defense in depth, not the guarantee.
	Snapshot(ctx context.Context) (Snapshot, error)
	// CurrentURL is used to read back a draft URL after saving.
	CurrentURL(ctx context.Context) (string, error)
}

// FieldRef identifies a form field a driver fills. Role is the semantic slot;
// Selector is the marketplace-specific hint declared in the driver manifest.
type FieldRef struct {
	Role     string // "title", "description", "price", "category", "condition", "photos"
	Selector string // marketplace-specific selector / accessible name from the manifest
}

// Control is one actionable control discovered on the page.
type Control struct {
	Role           string `json:"role"`            // "button", "link", "submit"
	AccessibleName string `json:"accessible_name"` // visible label / aria-label
	ID             string `json:"id"`              // stable handle for Click + click-log assertions
}

// Snapshot is an opaque pre-terminal-action capture handle (a11y tree ref or
// screenshot path).
type Snapshot struct {
	Kind string `json:"kind"` // "a11y_tree", "screenshot"
	Ref  string `json:"ref"`
}

// DraftLister is an optional capability a Page may implement so a driver can
// enumerate existing drafts for receipt reconciliation. Not every surface
// exposes a machine-readable drafts index, so it is kept off the core Page
// interface and discovered by type assertion. A driver's ListDrafts returns an
// empty list (with a warning) when the page does not implement it.
type DraftLister interface {
	Drafts(ctx context.Context) ([]DraftSummary, error)
}
