// Package driver is the open-source core of the grenier marketplace runner.
//
// It defines the capability interface every marketplace driver implements, the
// data shapes the night-operator Skill hands off, and the runner that owns the
// browser lifecycle, persistence, retries, audit log, and receipts.
//
// The product contract from the Skill is load-bearing here (DEC-0130):
//
//	save as draft is terminal. No driver may publish, send, message, delete, or
//	accept anything without explicit human validation after the wake-up handoff.
//
// So the V1 interface has exactly one write path — FillDraft. Publish,
// MessageBuyer, AcceptOffer, DeleteListing deliberately do not exist: you cannot
// call a method that isn't there.
package driver

import "context"

// Driver is the small capability interface every marketplace implements. It is
// intentionally not a marketplace-specific API wrapper — the runner talks to any
// marketplace through these six methods only.
type Driver interface {
	Marketplace() Marketplace
	Capabilities(ctx context.Context) (Capabilities, error)
	FillDraft(ctx context.Context, page Page, item DraftInput) (DraftResult, error)
	ListDrafts(ctx context.Context, page Page) ([]DraftSummary, error)
	LoadState(ctx context.Context) (DriverState, error)
	SaveState(ctx context.Context, state DriverState) error
}

// Marketplace identifies one marketplace surface.
type Marketplace struct {
	ID      string `json:"id"`      // "vinted-fr", "leboncoin-fr"
	Country string `json:"country"` // "FR"
	Name    string `json:"name"`
}

// PermissionPosture and detection risk are different axes: a driver can be
// low-detection and still ToS-forbidden, or high-detection and officially
// authorized. Keep them separate so review is honest.
type PermissionPosture string

const (
	PermissionOfficialAPI PermissionPosture = "official_api"
	PermissionAuthorized  PermissionPosture = "authorized"
	PermissionUserBrowser PermissionPosture = "user_browser_tos_hostile"
	PermissionForbidden   PermissionPosture = "forbidden"
	PermissionUnknown     PermissionPosture = "unknown"
)

// RiskLevel is the detection-risk axis, orthogonal to permission.
type RiskLevel string

const (
	RiskLow     RiskLevel = "low"
	RiskMedium  RiskLevel = "medium"
	RiskHigh    RiskLevel = "high"
	RiskUnknown RiskLevel = "unknown"
)

// DraftStop describes whether the marketplace exposes a stable draft stop. A
// marketplace with no stable draft stop is not eligible for a V1 driver — it can
// exist as a research note, not shipped code.
type DraftStop string

const (
	DraftStopNative  DraftStop = "native_save_draft" // an explicit "save as draft" control
	DraftStopNone    DraftStop = "none"              // publish-only: disqualified for V1
	DraftStopUnknown DraftStop = "unknown"
)

// Capabilities is what a driver can do, declared up front so the runner can gate
// a run before touching the browser.
type Capabilities struct {
	FillDraft      bool              `json:"fill_draft"`
	ListDrafts     bool              `json:"list_drafts"`
	UploadPhotos   bool              `json:"upload_photos"`
	DraftStop      DraftStop         `json:"draft_stop"`
	Permission     PermissionPosture `json:"permission"`
	DetectionRisk  RiskLevel         `json:"detection_risk"`
	MaxItemsPerRun int               `json:"max_items_per_run"`
}

// DraftInput is one listing the night operator prepared for a marketplace.
type DraftInput struct {
	ItemID      string            `json:"item_id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	PriceCents  int               `json:"price_cents"`
	Currency    string            `json:"currency"`
	Category    string            `json:"category"`
	Condition   string            `json:"condition"`
	PhotoPaths  []string          `json:"photo_paths"`
	Attributes  map[string]string `json:"attributes"`
}

// DraftStatus is the terminal outcome for one item. Every value is honest about
// what happened — the Skill must be able to resume without pretending the night
// succeeded.
type DraftStatus string

const (
	StatusDraftSaved       DraftStatus = "draft_saved"
	StatusDryRun           DraftStatus = "dry_run" // fields filled, terminal action deliberately not taken
	StatusSkippedAmbiguous DraftStatus = "skipped_ambiguous_submit"
	StatusSkippedPublish   DraftStatus = "skipped_publish_only"   // no draft stop, publish-only page: disqualified
	StatusSkippedDrift     DraftStatus = "skipped_selector_drift" // selectors did not match: skip, never best-effort publish
	StatusSkippedFields    DraftStatus = "skipped_missing_fields"
	StatusNeedsHuman       DraftStatus = "needs_human_review"
	StatusError            DraftStatus = "error"
)

// Skipped reports whether a status means the item was not drafted (for receipt
// completeness accounting).
func (s DraftStatus) Skipped() bool {
	switch s {
	case StatusSkippedAmbiguous, StatusSkippedPublish, StatusSkippedDrift,
		StatusSkippedFields, StatusNeedsHuman, StatusError:
		return true
	}
	return false
}

// DraftResult is the outcome of FillDraft for one item.
type DraftResult struct {
	ItemID        string      `json:"item_id"`
	Marketplace   string      `json:"marketplace"`
	Status        DraftStatus `json:"status"`
	DraftURL      string      `json:"draft_url,omitempty"`
	SkippedReason string      `json:"skipped_reason,omitempty"`
	Warnings      []string    `json:"warnings,omitempty"`
}

// DraftSummary is one existing draft, used to reconcile receipts on resume.
type DraftSummary struct {
	ItemID   string `json:"item_id"`
	DraftURL string `json:"draft_url"`
	Title    string `json:"title"`
}

// DriverState is opaque per-driver persisted state (session hints, cursors). The
// runner persists it between runs; the driver decides what goes in it.
type DriverState struct {
	Marketplace string            `json:"marketplace"`
	Values      map[string]string `json:"values,omitempty"`
}
