package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func writeTempCSV(t *testing.T, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "transactions.csv")
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("writing temp CSV: %v", err)
	}
	return path
}

// TestRunOpenMissingFileErrors doesn't depend on the exercise TODOs —
// os.Open fails immediately for a path that doesn't exist. Already
// passing.
func TestRunOpenMissingFileErrors(t *testing.T) {
	var out bytes.Buffer
	err := run(filepath.Join(t.TempDir(), "does-not-exist.csv"), "table", &out)
	if err == nil {
		t.Fatal("run() error = nil, want an error for a missing file")
	}
}

// TestRunRejectsUnknownFormat doesn't depend on the exercise TODOs —
// the format switch's default branch runs regardless of whether
// ParseTransactions/Summarize are implemented yet. Already passing.
func TestRunRejectsUnknownFormat(t *testing.T) {
	path := writeTempCSV(t, "date,category,amount\n")
	var out bytes.Buffer

	err := run(path, "xml", &out)
	if err == nil {
		t.Fatal("run() error = nil, want an error for an unknown format")
	}
}

// TestRunEndToEndJSON is the full pipeline: parse -> summarize ->
// format. Skipped until ParseTransactions and Summarize are both
// implemented (see csv.go, summary.go).
func TestRunEndToEndJSON(t *testing.T) {
	t.Skip("TODO(exercise): implement ParseTransactions and Summarize, then remove this Skip")

	path := writeTempCSV(t, "date,category,amount\n"+
		"2024-01-01,groceries,45.50\n"+
		"2024-01-02,transport,12.00\n")
	var out bytes.Buffer

	if err := run(path, "json", &out); err != nil {
		t.Fatalf("run() error = %v, want nil", err)
	}
	if !strings.Contains(out.String(), `"total": 57.5`) {
		t.Errorf("output = %s, want it to contain \"total\": 57.5", out.String())
	}
}

// TestRunEndToEndMalformedRowNeverPanics is the project's own
// completion criteria, verbatim: a malformed row produces a clear
// wrapped error, not a panic and not a silent skip. Skipped until
// ParseTransactions is implemented.
func TestRunEndToEndMalformedRowNeverPanics(t *testing.T) {
	t.Skip("TODO(exercise): implement ParseTransactions, then remove this Skip")

	path := writeTempCSV(t, "date,category,amount\n"+
		"2024-01-01,groceries,not-a-number\n")
	var out bytes.Buffer

	err := run(path, "table", &out)
	if err == nil {
		t.Fatal("run() error = nil, want an error for the malformed row")
	}
}
