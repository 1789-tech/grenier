package driver_test

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/1789-tech/grenier/driver"
	"github.com/1789-tech/grenier/driver/fixturemarket"
	"github.com/1789-tech/grenier/drivers/vintedfr"
)

func mustManifest(t *testing.T) driver.Manifest {
	t.Helper()
	m, err := vintedfr.Manifest()
	if err != nil {
		t.Fatalf("manifest: %v", err)
	}
	return m
}

func handoff(n int, dryRun bool) driver.Handoff {
	items := make([]driver.DraftInput, n)
	for i := range items {
		items[i] = driver.DraftInput{
			ItemID:     "grn-" + string(rune('a'+i)),
			Title:      "Item",
			PriceCents: 1000,
			Currency:   "EUR",
		}
	}
	return driver.Handoff{
		RunID:       "grn-run-test-001",
		Country:     "FR",
		Marketplace: "vinted-fr",
		DryRun:      dryRun,
		Items:       items,
	}
}

func happyPage() *fixturemarket.Market {
	return fixturemarket.New(fixturemarket.Scenario{
		Name: "happy",
		Controls: []driver.Control{
			{Role: "button", AccessibleName: "Publier", ID: "btn-publish"},
			{Role: "button", AccessibleName: "Enregistrer le brouillon", ID: "btn-draft"},
		},
		DraftURL: "https://www.vinted.fr/items/1/edit?draft=1",
	})
}

func TestRun_WritesIncrementalReceipt(t *testing.T) {
	dir := t.TempDir()
	d, _ := vintedfr.New()
	rec, err := driver.Run(context.Background(), driver.RunConfig{
		Handoff:     handoff(2, false),
		Driver:      d,
		Manifest:    mustManifest(t),
		Page:        happyPage(),
		ReceiptPath: filepath.Join(dir, "receipt.json"),
		AuditPath:   filepath.Join(dir, "audit.jsonl"),
	})
	if err != nil {
		t.Fatalf("Run: %v", err)
	}
	drafted, skipped, dry := rec.Summary()
	if drafted != 2 || skipped != 0 || dry != 0 {
		t.Fatalf("summary drafted=%d skipped=%d dry=%d", drafted, skipped, dry)
	}

	// Receipt is valid JSON on disk.
	data, err := os.ReadFile(filepath.Join(dir, "receipt.json"))
	if err != nil {
		t.Fatalf("read receipt: %v", err)
	}
	var onDisk driver.Receipt
	if err := json.Unmarshal(data, &onDisk); err != nil {
		t.Fatalf("receipt not valid json: %v", err)
	}
	if len(onDisk.Results) != 2 {
		t.Fatalf("on-disk results = %d", len(onDisk.Results))
	}

	// Audit log has content.
	if info, err := os.Stat(filepath.Join(dir, "audit.jsonl")); err != nil || info.Size() == 0 {
		t.Fatalf("audit log missing/empty: err=%v", err)
	}
}

func TestRun_ResumesWithoutRedoing(t *testing.T) {
	dir := t.TempDir()
	receiptPath := filepath.Join(dir, "receipt.json")

	// Pre-seed a receipt as if the first item already ran.
	seed := driver.Receipt{
		RunID:       "grn-run-test-001",
		Marketplace: "vinted-fr",
		Results:     []driver.DraftResult{{ItemID: "grn-a", Marketplace: "vinted-fr", Status: driver.StatusDraftSaved, DraftURL: "x"}},
	}
	data, _ := json.MarshalIndent(seed, "", "  ")
	if err := os.WriteFile(receiptPath, data, 0o644); err != nil {
		t.Fatal(err)
	}

	page := happyPage()
	d, _ := vintedfr.New()
	rec, err := driver.Run(context.Background(), driver.RunConfig{
		Handoff:     handoff(2, false),
		Driver:      d,
		Manifest:    mustManifest(t),
		Page:        page,
		ReceiptPath: receiptPath,
	})
	if err != nil {
		t.Fatalf("Run: %v", err)
	}
	if len(rec.Results) != 2 {
		t.Fatalf("results = %d, want 2 (1 resumed + 1 new)", len(rec.Results))
	}
	// Only the second item should have been processed this run: one draft click.
	if len(page.Clicked()) != 1 {
		t.Fatalf("expected exactly 1 click on resume, got %d (%v)", len(page.Clicked()), page.Clicked())
	}
}

func TestRun_DryRunPropagates(t *testing.T) {
	dir := t.TempDir()
	page := happyPage()
	d, _ := vintedfr.New()
	rec, err := driver.Run(context.Background(), driver.RunConfig{
		Handoff:     handoff(1, true),
		Driver:      d,
		Manifest:    mustManifest(t),
		Page:        page,
		ReceiptPath: filepath.Join(dir, "receipt.json"),
	})
	if err != nil {
		t.Fatalf("Run: %v", err)
	}
	_, _, dry := rec.Summary()
	if dry != 1 {
		t.Fatalf("dry = %d, want 1", dry)
	}
	if len(page.Clicked()) != 0 {
		t.Fatalf("dry-run run must not click, got %v", page.Clicked())
	}
}

func TestRun_EnforcesMaxItemsPerRunVisibly(t *testing.T) {
	dir := t.TempDir()
	d, _ := vintedfr.New()
	// Manifest cap is 6; hand off 8. Items 7 and 8 must appear as needs_human_review.
	rec, err := driver.Run(context.Background(), driver.RunConfig{
		Handoff:     handoff(8, false),
		Driver:      d,
		Manifest:    mustManifest(t),
		Page:        happyPage(),
		ReceiptPath: filepath.Join(dir, "receipt.json"),
	})
	if err != nil {
		t.Fatalf("Run: %v", err)
	}
	if len(rec.Results) != 8 {
		t.Fatalf("results = %d, want 8 (all items recorded)", len(rec.Results))
	}
	drafted, skipped, _ := rec.Summary()
	if drafted != 6 || skipped != 2 {
		t.Fatalf("drafted=%d skipped=%d, want 6/2", drafted, skipped)
	}
	for _, r := range rec.Results[6:] {
		if r.Status != driver.StatusNeedsHuman {
			t.Fatalf("over-cap item status = %q, want needs_human_review", r.Status)
		}
	}
}

func TestRun_RejectsMarketplaceMismatch(t *testing.T) {
	dir := t.TempDir()
	d, _ := vintedfr.New()
	h := handoff(1, false)
	h.Marketplace = "leboncoin-fr"
	_, err := driver.Run(context.Background(), driver.RunConfig{
		Handoff:     h,
		Driver:      d,
		Manifest:    mustManifest(t),
		Page:        happyPage(),
		ReceiptPath: filepath.Join(dir, "receipt.json"),
	})
	if err == nil {
		t.Fatalf("expected mismatch error")
	}
}
