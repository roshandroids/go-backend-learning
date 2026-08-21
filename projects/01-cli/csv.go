package main

import (
	"encoding/csv"
	"fmt"
	"io"
)

// ParseTransactions reads CSV rows shaped date,category,amount (with a
// header row to skip) and returns the parsed Transactions.
//
// TODO(exercise, Level 6 — Build): for each data row (skip row 0, the
// header):
//  1. Parse column 2 (amount) with strconv.ParseFloat(row[2], 64).
//  2. If parsing fails, return nil and a wrapped error identifying the
//     1-indexed row number and the underlying error — e.g.
//     fmt.Errorf("row %d: parsing amount %q: %w", rowNum, row[2], err).
//     Do NOT panic and do NOT silently skip the bad row; the whole call
//     must fail clearly so the caller knows the input was malformed.
//  3. Otherwise append a Transaction{Date: row[0], Category: row[1],
//     Amount: amount} to the result.
//
// Return an empty, non-nil slice (not an error) for a file with only a
// header and no data rows.
func ParseTransactions(r io.Reader) ([]Transaction, error) {
	reader := csv.NewReader(r)

	if _, err := reader.Read(); err != nil { // skip the header row
		return nil, fmt.Errorf("reading header: %w", err)
	}

	// TODO: replace this placeholder with the real parsing loop.
	return []Transaction{}, nil
}
