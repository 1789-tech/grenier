package driver

import (
	"context"
	"errors"
	"fmt"
)

// dryRunKey carries the dry-run flag through context so FillDraft honors it
// without widening the interface. Every driver must check IsDryRun and stop
// before the terminal draft action when it is set.
type dryRunKey struct{}

// WithDryRun marks ctx as a dry run.
func WithDryRun(ctx context.Context) context.Context {
	return context.WithValue(ctx, dryRunKey{}, true)
}

// IsDryRun reports whether ctx is a dry run. Drivers call this.
func IsDryRun(ctx context.Context) bool {
	v, _ := ctx.Value(dryRunKey{}).(bool)
	return v
}

// RunConfig is everything the runner needs for one handoff.
type RunConfig struct {
	Handoff     Handoff
	Driver      Driver
	Manifest    Manifest
	Page        Page
	ReceiptPath string
	AuditPath   string // optional; "" disables the audit log
	MaxRetries  int    // transient-error retries per item (default 0)
	DryRun      bool   // force dry run regardless of handoff
}

// Run executes one handoff: gates the run against the driver's capabilities,
// fills a draft per item (honoring dry-run), and writes an incremental receipt
// plus an audit log. It never publishes — publish is not in the interface, and
// the terminal action is chosen by the driver's manifest-declared draft matcher.
func Run(ctx context.Context, cfg RunConfig) (Receipt, error) {
	if cfg.Driver == nil {
		return Receipt{}, errors.New("run: nil driver")
	}
	if cfg.Page == nil {
		return Receipt{}, errors.New("run: nil page")
	}
	if err := cfg.Manifest.Validate(); err != nil {
		return Receipt{}, err
	}

	mkt := cfg.Driver.Marketplace()
	if cfg.Handoff.Marketplace != mkt.ID {
		return Receipt{}, fmt.Errorf("run: handoff marketplace %q != loaded driver %q",
			cfg.Handoff.Marketplace, mkt.ID)
	}

	caps, err := cfg.Driver.Capabilities(ctx)
	if err != nil {
		return Receipt{}, fmt.Errorf("run: capabilities: %w", err)
	}
	if !caps.FillDraft {
		return Receipt{}, fmt.Errorf("run: driver %q cannot fill drafts", mkt.ID)
	}
	if caps.DraftStop == DraftStopNone {
		return Receipt{}, fmt.Errorf("run: driver %q has no draft stop (publish-only) — disqualified for V1", mkt.ID)
	}

	dryRun := cfg.DryRun || cfg.Handoff.DryRun
	runCtx := ctx
	if dryRun {
		runCtx = WithDryRun(ctx)
	}

	rw, err := OpenReceipt(cfg.ReceiptPath, cfg.Handoff.RunID, mkt.ID, dryRun)
	if err != nil {
		return Receipt{}, err
	}

	var audit *AuditLog
	if cfg.AuditPath != "" {
		audit, err = OpenAuditLog(cfg.AuditPath)
		if err != nil {
			return Receipt{}, err
		}
		defer audit.Close()
	}
	audit.Write(AuditEntry{RunID: cfg.Handoff.RunID, Event: "run_start",
		Detail: fmt.Sprintf("driver=%s items=%d dry_run=%v", mkt.ID, len(cfg.Handoff.Items), dryRun)})

	done := rw.Done()
	for i, item := range cfg.Handoff.Items {
		if done[item.ItemID] {
			continue // resume: already recorded
		}

		// Enforce the manifest run cap, visibly, instead of silently drafting past it.
		if caps.MaxItemsPerRun > 0 && i >= caps.MaxItemsPerRun {
			res := DraftResult{
				ItemID:        item.ItemID,
				Marketplace:   mkt.ID,
				Status:        StatusNeedsHuman,
				SkippedReason: fmt.Sprintf("exceeds max_items_per_run=%d — left for a later run", caps.MaxItemsPerRun),
			}
			if err := rw.Append(res); err != nil {
				return rw.Receipt(), err
			}
			audit.Write(AuditEntry{RunID: cfg.Handoff.RunID, ItemID: item.ItemID, Event: "skipped", Detail: res.SkippedReason})
			continue
		}

		audit.Write(AuditEntry{RunID: cfg.Handoff.RunID, ItemID: item.ItemID, Event: "item_start"})
		res := fillWithRetries(runCtx, cfg, item, audit)
		res.ItemID = item.ItemID
		res.Marketplace = mkt.ID
		if err := rw.Append(res); err != nil {
			return rw.Receipt(), err
		}
		switch {
		case res.Status == StatusDraftSaved:
			audit.Write(AuditEntry{RunID: cfg.Handoff.RunID, ItemID: item.ItemID, Event: "draft_saved", Detail: res.DraftURL})
		case res.Status == StatusError:
			audit.Write(AuditEntry{RunID: cfg.Handoff.RunID, ItemID: item.ItemID, Event: "error", Detail: res.SkippedReason})
		default:
			audit.Write(AuditEntry{RunID: cfg.Handoff.RunID, ItemID: item.ItemID, Event: "skipped", Detail: string(res.Status) + ": " + res.SkippedReason})
		}
	}

	return rw.Receipt(), nil
}

// fillWithRetries retries only on transient errors (FillDraft returning a
// non-nil error). A skip is a deliberate, terminal safety decision — never
// retried into a best-effort publish.
func fillWithRetries(ctx context.Context, cfg RunConfig, item DraftInput, audit *AuditLog) DraftResult {
	var lastErr error
	for attempt := 0; attempt <= cfg.MaxRetries; attempt++ {
		res, err := cfg.Driver.FillDraft(ctx, cfg.Page, item)
		if err == nil {
			return res
		}
		lastErr = err
		audit.Write(AuditEntry{RunID: cfg.Handoff.RunID, ItemID: item.ItemID, Event: "error",
			Detail: fmt.Sprintf("attempt %d: %v", attempt+1, err)})
	}
	return DraftResult{
		Status:        StatusError,
		SkippedReason: fmt.Sprintf("fill_draft failed after %d attempt(s): %v", cfg.MaxRetries+1, lastErr),
	}
}
