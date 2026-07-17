// Package fixturemarket is an in-memory fake marketplace DOM that implements
// driver.Page. It lets a real driver be exercised end-to-end — fill fields,
// discover controls, take the terminal draft action — without launching a
// browser or touching a real account. It records every click so tests can
// assert a publish button is present yet never clicked.
//
// This is how the V1 reference driver is proven while the ToS posture of a real
// marketplace account is still an open Manfred decision: the exact same driver
// code runs against this fixture in CI and against a real user-owned browser in
// production, because both satisfy driver.Page.
package fixturemarket

import (
	"context"
	"fmt"
	"strings"

	"github.com/1789-tech/grenier/driver"
)

// Scenario configures which DOM the fixture presents.
type Scenario struct {
	Name string
	// Controls presented on the listing form. Include a publish button here to
	// exercise the "publish present, never clicked" assertion.
	Controls []driver.Control
	// DraftURL is what CurrentURL returns after the draft control is clicked.
	DraftURL string
	// FailUploadOnce makes the first UploadPhotos call return a transient error,
	// to exercise runner retries.
	FailUploadOnce bool
	// ExistingDrafts is what Drafts returns, for ListDrafts reconciliation.
	ExistingDrafts []driver.DraftSummary
}

// Market is a fake marketplace page. The zero value is not usable; use New.
type Market struct {
	scn        Scenario
	fields     map[string]string
	photos     []string
	clicks     []driver.Control
	url        string
	uploadFail bool
}

// New builds a Market for a scenario.
func New(scn Scenario) *Market {
	return &Market{
		scn:        scn,
		fields:     map[string]string{},
		url:        "https://fixture.invalid/items/new",
		uploadFail: scn.FailUploadOnce,
	}
}

// --- driver.Page implementation ---------------------------------------------

func (m *Market) Navigate(_ context.Context, url string) error {
	m.url = url
	return nil
}

func (m *Market) UploadPhotos(_ context.Context, _ driver.FieldRef, paths []string) error {
	if m.uploadFail {
		m.uploadFail = false
		return fmt.Errorf("fixture: transient upload failure")
	}
	m.photos = append(m.photos, paths...)
	return nil
}

func (m *Market) Fill(_ context.Context, field driver.FieldRef, value string) error {
	m.fields[field.Role] = value
	return nil
}

func (m *Market) Controls(_ context.Context) ([]driver.Control, error) {
	out := make([]driver.Control, len(m.scn.Controls))
	copy(out, m.scn.Controls)
	return out, nil
}

func (m *Market) Click(_ context.Context, c driver.Control) error {
	m.clicks = append(m.clicks, c)
	// Clicking the draft control navigates to the saved-draft URL.
	if m.scn.DraftURL != "" && strings.Contains(strings.ToLower(c.AccessibleName), "brouillon") {
		m.url = m.scn.DraftURL
	}
	return nil
}

func (m *Market) Snapshot(_ context.Context) (driver.Snapshot, error) {
	return driver.Snapshot{Kind: "a11y_tree", Ref: "fixture://snapshot/" + m.scn.Name}, nil
}

func (m *Market) CurrentURL(_ context.Context) (string, error) {
	return m.url, nil
}

// Drafts implements the optional driver.DraftLister capability.
func (m *Market) Drafts(_ context.Context) ([]driver.DraftSummary, error) {
	out := make([]driver.DraftSummary, len(m.scn.ExistingDrafts))
	copy(out, m.scn.ExistingDrafts)
	return out, nil
}

// --- test introspection ------------------------------------------------------

// Clicked returns the accessible names of every control that was clicked, in
// order — the ground truth for anti-publish assertions.
func (m *Market) Clicked() []string {
	names := make([]string, len(m.clicks))
	for i, c := range m.clicks {
		names[i] = c.AccessibleName
	}
	return names
}

// ClickedID reports whether a control with the given ID was ever clicked.
func (m *Market) ClickedID(id string) bool {
	for _, c := range m.clicks {
		if c.ID == id {
			return true
		}
	}
	return false
}

// Field returns the value filled for a role.
func (m *Market) Field(role string) string { return m.fields[role] }

// Photos returns the uploaded photo paths.
func (m *Market) Photos() []string { return m.photos }
