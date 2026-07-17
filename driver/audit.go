package driver

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

// AuditEntry is one append-only line in the run's audit log. The log is defense
// in depth — a readable trail of what the runner did per item — not the safety
// guarantee itself (that rests on driver behavior, tests, and receipts).
type AuditEntry struct {
	RunID       string   `json:"run_id"`
	ItemID      string   `json:"item_id,omitempty"`
	Event       string   `json:"event"` // "run_start", "item_start", "fields_filled", "control_chosen", "snapshot", "draft_saved", "skipped", "error"
	Detail      string   `json:"detail,omitempty"`
	Control     string   `json:"control,omitempty"`
	Snapshot    string   `json:"snapshot,omitempty"`
	SeenControl []string `json:"seen_controls,omitempty"`
}

// AuditLog is a JSONL writer. A nil *AuditLog is a valid no-op sink.
type AuditLog struct {
	w io.WriteCloser
}

// OpenAuditLog opens (append mode) a JSONL audit log at path.
func OpenAuditLog(path string) (*AuditLog, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return nil, err
	}
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		return nil, err
	}
	return &AuditLog{w: f}, nil
}

// Write appends one entry. Safe to call on a nil log (no-op).
func (a *AuditLog) Write(e AuditEntry) {
	if a == nil || a.w == nil {
		return
	}
	data, err := json.Marshal(e)
	if err != nil {
		return
	}
	_, _ = a.w.Write(append(data, '\n'))
}

// Close closes the underlying file. Safe on nil.
func (a *AuditLog) Close() error {
	if a == nil || a.w == nil {
		return nil
	}
	return a.w.Close()
}
