// Package vintedfr is the V1 reference full-browser driver for Vinted France.
//
// IMPORTANT — posture: DEC-0130 leaves acceptance of the
// `user_browser_tos_hostile` posture for the first driver as a Manfred
// arbitrage. Until that is ruled, this driver is only ever run against the
// fixturemarket DOM (dry-run / CI). The exact same code path runs against a real
// user-owned browser once the posture is accepted, because both satisfy
// driver.Page — but this package never presumes that acceptance.
//
// It proves the FillDraft path end-to-end: navigate, upload photos, fill
// title/description/price/category/condition, and take the terminal
// save-as-draft action only — never publish. The publish control is discovered
// solely so it can be avoided.
package vintedfr

import (
	"context"
	_ "embed"
	"fmt"
	"strings"

	"github.com/1789-tech/grenier/driver"
)

//go:embed manifest.toml
var manifestTOML []byte

// Manifest parses and validates the embedded driver manifest.
func Manifest() (driver.Manifest, error) {
	m, err := driver.ParseManifest(manifestTOML)
	if err != nil {
		return driver.Manifest{}, err
	}
	if err := m.Validate(); err != nil {
		return driver.Manifest{}, err
	}
	return m, nil
}

func init() {
	driver.Register("vinted-fr", func(m driver.Manifest) (driver.Driver, error) {
		return &Driver{m: m}, nil
	})
}

// Driver implements driver.Driver for Vinted FR.
type Driver struct {
	m     driver.Manifest
	state driver.DriverState
}

// New builds a driver from the embedded manifest. Convenience for callers not
// going through the registry.
func New() (*Driver, error) {
	m, err := Manifest()
	if err != nil {
		return nil, err
	}
	return &Driver{m: m}, nil
}

func (d *Driver) Marketplace() driver.Marketplace {
	return driver.Marketplace{ID: d.m.ID, Country: d.m.Country, Name: d.m.DisplayName}
}

func (d *Driver) Capabilities(context.Context) (driver.Capabilities, error) {
	return driver.Capabilities{
		FillDraft:      true,
		ListDrafts:     true,
		UploadPhotos:   true,
		DraftStop:      driver.DraftStopNative,
		Permission:     d.m.PermissionPosture,
		DetectionRisk:  d.m.DetectionRisk,
		MaxItemsPerRun: d.m.MaxItemsPerRun,
	}, nil
}

// FillDraft fills one listing up to — and only up to — the save-as-draft action.
func (d *Driver) FillDraft(ctx context.Context, page driver.Page, item driver.DraftInput) (driver.DraftResult, error) {
	res := driver.DraftResult{ItemID: item.ItemID, Marketplace: d.m.ID}

	// Fail closed on missing required fields rather than drafting a broken listing.
	if strings.TrimSpace(item.Title) == "" || item.PriceCents <= 0 {
		res.Status = driver.StatusSkippedFields
		res.SkippedReason = "missing required title or price"
		return res, nil
	}

	if err := page.Navigate(ctx, d.m.NewListingURL); err != nil {
		return driver.DraftResult{}, fmt.Errorf("navigate: %w", err)
	}

	if len(item.PhotoPaths) > 0 {
		if err := page.UploadPhotos(ctx, driver.FieldRef{Role: "photos", Selector: "input[type=file]"}, item.PhotoPaths); err != nil {
			return driver.DraftResult{}, fmt.Errorf("upload photos: %w", err)
		}
	}

	fills := []struct {
		role, selector, value string
	}{
		{"title", "#title", item.Title},
		{"description", "#description", item.Description},
		{"price", "#price", formatPriceEUR(item.PriceCents, item.Currency)},
		{"category", "#category", item.Category},
		{"condition", "#status", item.Condition},
	}
	for _, f := range fills {
		if f.value == "" {
			continue
		}
		if err := page.Fill(ctx, driver.FieldRef{Role: f.role, Selector: f.selector}, f.value); err != nil {
			return driver.DraftResult{}, fmt.Errorf("fill %s: %w", f.role, err)
		}
	}
	for k, v := range item.Attributes {
		if err := page.Fill(ctx, driver.FieldRef{Role: "attr:" + k, Selector: "[data-attr=" + k + "]"}, v); err != nil {
			return driver.DraftResult{}, fmt.Errorf("fill attr %s: %w", k, err)
		}
	}

	// Resolve the terminal control via the shared, tested safety helper. It never
	// returns a publish control; a non-empty Status means "do not click, skip".
	controls, err := page.Controls(ctx)
	if err != nil {
		return driver.DraftResult{}, fmt.Errorf("controls: %w", err)
	}
	choice := driver.SelectTerminalControl(controls, d.m)
	if choice.Status != "" {
		res.Status = choice.Status
		res.SkippedReason = choice.Reason
		return res, nil
	}

	// Defense in depth: capture a pre-action snapshot for the audit trail.
	if _, err := page.Snapshot(ctx); err != nil {
		res.Warnings = append(res.Warnings, "snapshot failed: "+err.Error())
	}

	// Dry run stops here: fields are filled, terminal action deliberately skipped.
	if driver.IsDryRun(ctx) {
		res.Status = driver.StatusDryRun
		res.SkippedReason = "dry-run: fields filled, save-as-draft not taken"
		return res, nil
	}

	// Belt-and-suspenders: never click a control whose name reads as publish,
	// even if the selector helper somehow returned one.
	if isPublishName(choice.Control.AccessibleName, d.m.PublishActionNames) {
		res.Status = driver.StatusSkippedAmbiguous
		res.SkippedReason = "resolved control matched a publish name — refusing to click"
		return res, nil
	}

	if err := page.Click(ctx, choice.Control); err != nil {
		return driver.DraftResult{}, fmt.Errorf("click draft control: %w", err)
	}

	url, err := page.CurrentURL(ctx)
	if err != nil {
		res.Warnings = append(res.Warnings, "could not read draft url: "+err.Error())
	}
	res.Status = driver.StatusDraftSaved
	res.DraftURL = url
	return res, nil
}

// ListDrafts reconciles receipts against existing drafts when the page exposes
// the optional DraftLister capability; otherwise returns an empty list.
func (d *Driver) ListDrafts(ctx context.Context, page driver.Page) ([]driver.DraftSummary, error) {
	if l, ok := page.(driver.DraftLister); ok {
		return l.Drafts(ctx)
	}
	return []driver.DraftSummary{}, nil
}

func (d *Driver) LoadState(context.Context) (driver.DriverState, error) { return d.state, nil }

func (d *Driver) SaveState(_ context.Context, state driver.DriverState) error {
	d.state = state
	return nil
}

func isPublishName(name string, publish []string) bool {
	name = strings.ToLower(strings.TrimSpace(name))
	for _, p := range publish {
		if p != "" && strings.Contains(name, p) {
			return true
		}
	}
	return false
}

// formatPriceEUR renders cents as a FR-style price string, e.g. 2800 -> "28,00 €".
func formatPriceEUR(cents int, currency string) string {
	if currency == "" {
		currency = "EUR"
	}
	sym := "€"
	if currency != "EUR" {
		sym = currency
	}
	return fmt.Sprintf("%d,%02d %s", cents/100, cents%100, sym)
}
