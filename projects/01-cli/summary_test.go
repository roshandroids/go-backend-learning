package main

import "testing"

// TestSummarize is the exercise: totals and per-category breakdown.
// Skipped until the TODO in summary.go is filled in.
func TestSummarize(t *testing.T) {
	t.Skip("TODO(exercise): implement Summarize, then remove this Skip")

	txs := []Transaction{
		{Date: "2024-01-01", Category: "groceries", Amount: 45.50},
		{Date: "2024-01-02", Category: "transport", Amount: 12.00},
		{Date: "2024-01-03", Category: "groceries", Amount: 10.00},
	}

	got := Summarize(txs)

	if got.Total != 67.50 {
		t.Errorf("Total = %v, want 67.5", got.Total)
	}
	if got.ByCategory["groceries"] != 55.50 {
		t.Errorf("ByCategory[groceries] = %v, want 55.5", got.ByCategory["groceries"])
	}
	if got.ByCategory["transport"] != 12.00 {
		t.Errorf("ByCategory[transport] = %v, want 12", got.ByCategory["transport"])
	}
}

// TestSummarizeEmpty proves an empty input doesn't panic on a nil map
// write (concepts/02-types-and-structs' lesson, resurfacing here).
// Skipped until the TODO in summary.go is filled in.
func TestSummarizeEmpty(t *testing.T) {
	t.Skip("TODO(exercise): implement Summarize, then remove this Skip")

	got := Summarize(nil)
	if got.Total != 0 {
		t.Errorf("Total = %v, want 0", got.Total)
	}
}
