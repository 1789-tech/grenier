package vintedfr

import (
	"context"
	"testing"

	"github.com/1789-tech/grenier/driver"
	"github.com/1789-tech/grenier/driver/fixturemarket"
)

// controls for the reference scenarios. The publish button is present in the
// happy path precisely so we can prove it is never clicked.
var (
	publishBtn = driver.Control{Role: "button", AccessibleName: "Publier", ID: "btn-publish"}
	draftBtn   = driver.Control{Role: "button", AccessibleName: "Enregistrer le brouillon", ID: "btn-draft"}
	draftBtn2  = driver.Control{Role: "button", AccessibleName: "Sauvegarder le brouillon", ID: "btn-draft-2"}
)

func sampleItem() driver.DraftInput {
	return driver.DraftInput{
		ItemID:      "grn-20260717-0001",
		Title:       "Veste denim Levi's taille M",
		Description: "Bon etat, coupe droite, a retirer a Paris.",
		PriceCents:  2800,
		Currency:    "EUR",
		Category:    "mode",
		Condition:   "good",
		PhotoPaths:  []string{"photos/grn-20260717-0001-1.jpg"},
		Attributes:  map[string]string{"brand": "Levi's", "size": "M"},
	}
}

func newDriver(t *testing.T) *Driver {
	t.Helper()
	d, err := New()
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	return d
}

// Happy path: a draft is saved, fields are filled, and — critically — the
// publish button that is present on the page is NEVER clicked.
func TestFillDraft_HappyPath_SavesDraftNeverPublishes(t *testing.T) {
	d := newDriver(t)
	mkt := fixturemarket.New(fixturemarket.Scenario{
		Name:     "happy",
		Controls: []driver.Control{publishBtn, draftBtn},
		DraftURL: "https://www.vinted.fr/items/12345/edit?draft=1",
	})

	res, err := d.FillDraft(context.Background(), mkt, sampleItem())
	if err != nil {
		t.Fatalf("FillDraft: %v", err)
	}
	if res.Status != driver.StatusDraftSaved {
		t.Fatalf("status = %q, want draft_saved", res.Status)
	}
	if res.DraftURL == "" {
		t.Fatalf("expected a draft URL")
	}

	// The core safety assertion the contract requires.
	if mkt.ClickedID("btn-publish") {
		t.Fatalf("PUBLISH BUTTON WAS CLICKED — contract violation")
	}
	if !mkt.ClickedID("btn-draft") {
		t.Fatalf("draft button was not clicked")
	}

	// Fields were actually filled.
	if got := mkt.Field("title"); got != "Veste denim Levi's taille M" {
		t.Errorf("title = %q", got)
	}
	if got := mkt.Field("price"); got != "28,00 €" {
		t.Errorf("price = %q, want 28,00 €", got)
	}
	if len(mkt.Photos()) != 1 {
		t.Errorf("photos = %v", mkt.Photos())
	}
}

// Ambiguous submit: two draft-matching controls -> skipped, nothing clicked.
func TestFillDraft_AmbiguousSubmit_Skips(t *testing.T) {
	d := newDriver(t)
	mkt := fixturemarket.New(fixturemarket.Scenario{
		Name:     "ambiguous",
		Controls: []driver.Control{publishBtn, draftBtn, draftBtn2},
	})

	res, err := d.FillDraft(context.Background(), mkt, sampleItem())
	if err != nil {
		t.Fatalf("FillDraft: %v", err)
	}
	if res.Status != driver.StatusSkippedAmbiguous {
		t.Fatalf("status = %q, want skipped_ambiguous_submit", res.Status)
	}
	if len(mkt.Clicked()) != 0 {
		t.Fatalf("nothing should be clicked on ambiguous submit, got %v", mkt.Clicked())
	}
}

// Publish-only page: publish present, no draft stop -> disqualified, nothing clicked.
func TestFillDraft_PublishOnly_Disqualified(t *testing.T) {
	d := newDriver(t)
	mkt := fixturemarket.New(fixturemarket.Scenario{
		Name:     "publish-only",
		Controls: []driver.Control{publishBtn},
	})

	res, err := d.FillDraft(context.Background(), mkt, sampleItem())
	if err != nil {
		t.Fatalf("FillDraft: %v", err)
	}
	if res.Status != driver.StatusSkippedPublish {
		t.Fatalf("status = %q, want skipped_publish_only", res.Status)
	}
	if mkt.ClickedID("btn-publish") {
		t.Fatalf("PUBLISH BUTTON WAS CLICKED on a publish-only page — contract violation")
	}
}

// Selector drift: neither draft nor publish matches -> skip, never best-effort.
func TestFillDraft_SelectorDrift_Skips(t *testing.T) {
	d := newDriver(t)
	mkt := fixturemarket.New(fixturemarket.Scenario{
		Name: "drift",
		Controls: []driver.Control{
			{Role: "button", AccessibleName: "Continuer", ID: "btn-next"},
			{Role: "button", AccessibleName: "Annuler", ID: "btn-cancel"},
		},
	})

	res, err := d.FillDraft(context.Background(), mkt, sampleItem())
	if err != nil {
		t.Fatalf("FillDraft: %v", err)
	}
	if res.Status != driver.StatusSkippedDrift {
		t.Fatalf("status = %q, want skipped_selector_drift", res.Status)
	}
	if len(mkt.Clicked()) != 0 {
		t.Fatalf("nothing should be clicked on selector drift, got %v", mkt.Clicked())
	}
}

// Dry-run: fields filled, terminal action deliberately not taken.
func TestFillDraft_DryRun_FillsButDoesNotClick(t *testing.T) {
	d := newDriver(t)
	mkt := fixturemarket.New(fixturemarket.Scenario{
		Name:     "dryrun",
		Controls: []driver.Control{publishBtn, draftBtn},
		DraftURL: "https://www.vinted.fr/items/12345/edit?draft=1",
	})

	ctx := driver.WithDryRun(context.Background())
	res, err := d.FillDraft(ctx, mkt, sampleItem())
	if err != nil {
		t.Fatalf("FillDraft: %v", err)
	}
	if res.Status != driver.StatusDryRun {
		t.Fatalf("status = %q, want dry_run", res.Status)
	}
	if len(mkt.Clicked()) != 0 {
		t.Fatalf("dry-run must not click anything, got %v", mkt.Clicked())
	}
	if mkt.Field("title") == "" {
		t.Fatalf("dry-run should still fill fields")
	}
}

// Missing required fields: fail closed.
func TestFillDraft_MissingFields_Skips(t *testing.T) {
	d := newDriver(t)
	mkt := fixturemarket.New(fixturemarket.Scenario{
		Name:     "missing",
		Controls: []driver.Control{publishBtn, draftBtn},
	})
	item := sampleItem()
	item.Title = ""

	res, err := d.FillDraft(context.Background(), mkt, item)
	if err != nil {
		t.Fatalf("FillDraft: %v", err)
	}
	if res.Status != driver.StatusSkippedFields {
		t.Fatalf("status = %q, want skipped_missing_fields", res.Status)
	}
	if len(mkt.Clicked()) != 0 {
		t.Fatalf("nothing clicked when fields missing, got %v", mkt.Clicked())
	}
}

func TestListDrafts_Reconciles(t *testing.T) {
	d := newDriver(t)
	want := []driver.DraftSummary{{ItemID: "grn-1", DraftURL: "https://x/1", Title: "A"}}
	mkt := fixturemarket.New(fixturemarket.Scenario{Name: "drafts", ExistingDrafts: want})

	got, err := d.ListDrafts(context.Background(), mkt)
	if err != nil {
		t.Fatalf("ListDrafts: %v", err)
	}
	if len(got) != 1 || got[0].ItemID != "grn-1" {
		t.Fatalf("ListDrafts = %v", got)
	}
}

func TestManifest_ValidAndNeutralCore(t *testing.T) {
	m, err := Manifest()
	if err != nil {
		t.Fatalf("Manifest: %v", err)
	}
	if m.PublishActionWired {
		t.Fatalf("publish_action_wired must be false")
	}
	if m.PermissionPosture != driver.PermissionUserBrowser {
		t.Fatalf("posture = %q", m.PermissionPosture)
	}
	if len(m.DraftActionNames) == 0 || len(m.PublishActionNames) == 0 {
		t.Fatalf("manifest must declare draft and publish matchers")
	}
}
