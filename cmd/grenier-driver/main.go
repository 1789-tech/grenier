// Command grenier-driver is the local runner CLI for the marketplace driver
// core (DEC-0130 / BRIEF-118). It reads a night-operator handoff.json, loads one
// marketplace driver by manifest ID, fills a draft per item, and writes an
// incremental receipt.json the grenier Skill reads on wake-up.
//
// It never publishes: the terminal action is save-as-draft, chosen by the
// driver's manifest-declared matcher, and publish is not part of the driver
// interface.
//
// Browser backends:
//
//   - fixture (default): drives the in-memory fixturemarket DOM. This is the
//     honest V1 proof — it exercises the full FillDraft path end-to-end without
//     touching a real marketplace account, which is required while the
//     `user_browser_tos_hostile` posture is still an open Manfred decision.
//   - real: not wired in V1. Selecting it errors, on purpose — wiring a real
//     user-owned browser session against a live account waits on the posture
//     ruling. The driver code is already browser-agnostic (driver.Page), so this
//     is a one-file addition once the posture is accepted.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/1789-tech/grenier/driver"
	"github.com/1789-tech/grenier/driver/fixturemarket"
	"github.com/1789-tech/grenier/drivers/vintedfr"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "grenier-driver:", err)
		os.Exit(1)
	}
}

func run() error {
	var (
		handoffPath = flag.String("handoff", "", "path to handoff.json (required)")
		receiptPath = flag.String("receipt", "", "path to receipt.json (default: alongside handoff)")
		auditPath   = flag.String("audit", "", "path to audit.jsonl (default: alongside receipt)")
		browser     = flag.String("browser", "fixture", "browser backend: fixture | real")
		dryRun      = flag.Bool("dry-run", false, "fill fields but never take the terminal draft action")
		maxRetries  = flag.Int("retries", 1, "transient-error retries per item")
	)
	flag.Parse()

	if *handoffPath == "" {
		return fmt.Errorf("--handoff is required (registered drivers: %v)", driver.Registered())
	}

	h, err := driver.LoadHandoff(*handoffPath)
	if err != nil {
		return err
	}

	recPath := *receiptPath
	if recPath == "" {
		recPath = filepath.Join(filepath.Dir(*handoffPath), "receipt.json")
	}
	audPath := *auditPath
	if audPath == "" {
		audPath = filepath.Join(filepath.Dir(recPath), "audit.jsonl")
	}

	// V1 ships one reference driver. Loading is by manifest ID from the handoff.
	m, err := vintedfr.Manifest()
	if err != nil {
		return err
	}
	if h.Marketplace != m.ID {
		return fmt.Errorf("handoff marketplace %q has no registered driver (V1 ships %q only)", h.Marketplace, m.ID)
	}
	drv, err := driver.Load(m)
	if err != nil {
		return err
	}

	page, err := openBrowser(*browser, m)
	if err != nil {
		return err
	}

	rec, err := driver.Run(context.Background(), driver.RunConfig{
		Handoff:     h,
		Driver:      drv,
		Manifest:    m,
		Page:        page,
		ReceiptPath: recPath,
		AuditPath:   audPath,
		MaxRetries:  *maxRetries,
		DryRun:      *dryRun,
	})
	if err != nil {
		return err
	}

	drafted, skipped, dry := rec.Summary()
	fmt.Printf("run %s (%s) via %s browser\n", rec.RunID, m.ID, *browser)
	fmt.Printf("  drafts saved:      %d\n", drafted)
	fmt.Printf("  dry-run filled:    %d\n", dry)
	fmt.Printf("  needs your eyes:   %d\n", skipped)
	fmt.Printf("  receipt:           %s\n", recPath)
	fmt.Printf("  audit log:         %s\n", audPath)
	return nil
}

// openBrowser builds a driver.Page for the chosen backend. Only the fixture is
// wired in V1; the real backend is deliberately gated on the ToS posture ruling.
func openBrowser(kind string, m driver.Manifest) (driver.Page, error) {
	switch kind {
	case "fixture":
		// A representative FR listing form: a publish button AND a draft stop.
		// The runner will fill fields and take only the draft action.
		return fixturemarket.New(fixturemarket.Scenario{
			Name: "cli-fixture",
			Controls: []driver.Control{
				{Role: "button", AccessibleName: "Publier", ID: "btn-publish"},
				{Role: "button", AccessibleName: "Enregistrer le brouillon", ID: "btn-draft"},
			},
			DraftURL: "https://www.vinted.fr/items/000000/edit?draft=1",
		}), nil
	case "real":
		return nil, fmt.Errorf("real browser backend is not wired in V1: the %q posture (%s) is an open Manfred decision — run with --browser=fixture until it is ruled",
			m.ID, m.PermissionPosture)
	default:
		return nil, fmt.Errorf("unknown browser backend %q (want: fixture | real)", kind)
	}
}
