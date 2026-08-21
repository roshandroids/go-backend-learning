package main

import (
	"strings"
	"testing"
)

// TestParseTransactionsValid is the exercise's happy path. Skipped until
// the TODO in csv.go is filled in.
func TestParseTransactionsValid(t *testing.T) {
	t.Skip("TODO(exercise): implement ParseTransactions, then remove this Skip")

	input := "date,category,amount\n" +
		"2024-01-01,groceries,45.50\n" +
		"2024-01-02,transport,12.00\n"

	txs, err := ParseTransactions(strings.NewReader(input))
	if err != nil {
		t.Fatalf("ParseTransactions() error = %v, want nil", err)
	}
	if len(txs) != 2 {
		t.Fatalf("got %d transactions, want 2", len(txs))
	}
	if txs[0].Category != "groceries" || txs[0].Amount != 45.50 {
		t.Errorf("txs[0] = %+v, want {2024-01-01 groceries 45.5}", txs[0])
	}
}

// TestParseTransactionsMalformedAmount is the exercise's core
// requirement: a bad row must produce a clear wrapped error, never a
// panic and never a silent skip. Skipped until the TODO is filled in.
func TestParseTransactionsMalformedAmount(t *testing.T) {
	t.Skip("TODO(exercise): implement ParseTransactions, then remove this Skip")

	input := "date,category,amount\n" +
		"2024-01-01,groceries,45.50\n" +
		"2024-01-02,transport,not-a-number\n"

	_, err := ParseTransactions(strings.NewReader(input))
	if err == nil {
		t.Fatal("ParseTransactions() error = nil, want an error for the malformed row")
	}
	if !strings.Contains(err.Error(), "row 2") {
		t.Errorf("error = %q, want it to mention the row number (row 2)", err.Error())
	}
}

// TestParseTransactionsHeaderOnly proves a header-only file returns an
// empty, non-nil slice rather than an error. Skipped until the TODO is
// filled in.
func TestParseTransactionsHeaderOnly(t *testing.T) {
	t.Skip("TODO(exercise): implement ParseTransactions, then remove this Skip")

	txs, err := ParseTransactions(strings.NewReader("date,category,amount\n"))
	if err != nil {
		t.Fatalf("ParseTransactions() error = %v, want nil", err)
	}
	if len(txs) != 0 {
		t.Errorf("got %d transactions, want 0", len(txs))
	}
}
