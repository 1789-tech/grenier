package driver

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Handoff is the file the night-operator Skill produces: the listings prepared
// for one marketplace run. The runner reads it, the Skill never drives the
// browser itself.
type Handoff struct {
	RunID       string       `json:"run_id"`
	Country     string       `json:"country"`
	Marketplace string       `json:"marketplace"`
	DryRun      bool         `json:"dry_run,omitempty"`
	Items       []DraftInput `json:"items"`
}

// LoadHandoff reads and validates a handoff.json.
func LoadHandoff(path string) (Handoff, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Handoff{}, err
	}
	var h Handoff
	if err := json.Unmarshal(data, &h); err != nil {
		return Handoff{}, fmt.Errorf("handoff %s: %w", path, err)
	}
	if h.RunID == "" {
		return Handoff{}, fmt.Errorf("handoff %s: run_id is required", path)
	}
	if h.Marketplace == "" {
		return Handoff{}, fmt.Errorf("handoff %s: marketplace is required", path)
	}
	return h, nil
}

// Receipt is what the Skill reads on wake-up. It is written incrementally — one
// result appended after every item, including skipped items — so a crash or
// re-run resumes from the file instead of pretending the night succeeded.
type Receipt struct {
	RunID       string        `json:"run_id"`
	Marketplace string        `json:"marketplace"`
	DryRun      bool          `json:"dry_run,omitempty"`
	Results     []DraftResult `json:"results"`
}

// Done reports item IDs already recorded, so a resumed run can skip them.
func (r Receipt) Done() map[string]bool {
	done := map[string]bool{}
	for _, res := range r.Results {
		done[res.ItemID] = true
	}
	return done
}

// Summary counts outcomes for the honest completeness report the Skill shows on
// wake-up.
func (r Receipt) Summary() (drafted, skipped, dryRun int) {
	for _, res := range r.Results {
		switch {
		case res.Status == StatusDraftSaved:
			drafted++
		case res.Status == StatusDryRun:
			dryRun++
		case res.Status.Skipped():
			skipped++
		}
	}
	return
}

// ReceiptWriter persists a receipt incrementally and atomically. Every Append
// rewrites the whole file via a temp-file rename, so the receipt on disk is
// always valid JSON even if the process dies mid-run.
type ReceiptWriter struct {
	path    string
	receipt Receipt
}

// OpenReceipt opens (or resumes) a receipt at path. If the file exists its prior
// results are loaded so the run resumes instead of duplicating work.
func OpenReceipt(path, runID, marketplace string, dryRun bool) (*ReceiptWriter, error) {
	rw := &ReceiptWriter{
		path: path,
		receipt: Receipt{
			RunID:       runID,
			Marketplace: marketplace,
			DryRun:      dryRun,
			Results:     []DraftResult{},
		},
	}
	if data, err := os.ReadFile(path); err == nil {
		var prior Receipt
		if err := json.Unmarshal(data, &prior); err != nil {
			return nil, fmt.Errorf("resume receipt %s: %w", path, err)
		}
		if prior.RunID == runID {
			rw.receipt = prior
		}
	}
	return rw, nil
}

// Append records one result and flushes the whole receipt to disk atomically.
func (rw *ReceiptWriter) Append(res DraftResult) error {
	rw.receipt.Results = append(rw.receipt.Results, res)
	return rw.flush()
}

// Receipt returns the current in-memory receipt.
func (rw *ReceiptWriter) Receipt() Receipt { return rw.receipt }

// Done exposes already-recorded item IDs for resume.
func (rw *ReceiptWriter) Done() map[string]bool { return rw.receipt.Done() }

func (rw *ReceiptWriter) flush() error {
	data, err := json.MarshalIndent(rw.receipt, "", "  ")
	if err != nil {
		return err
	}
	tmp := rw.path + ".tmp"
	if err := os.MkdirAll(filepath.Dir(rw.path), 0o755); err != nil {
		return err
	}
	if err := os.WriteFile(tmp, append(data, '\n'), 0o644); err != nil {
		return err
	}
	return os.Rename(tmp, rw.path)
}
