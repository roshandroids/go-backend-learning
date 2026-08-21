package main

import (
	"encoding/json"
	"fmt"
	"io"
	"sort"
)

// summaryJSON is the JSON shape for Summary — a plain struct with
// exported fields is enough; no need for a separate DTO type at this
// size.
type summaryJSON struct {
	Total      float64            `json:"total"`
	ByCategory map[string]float64 `json:"by_category"`
}

// WriteJSON encodes s as JSON. Already implemented — JSON encoding is a
// "skim" topic, conceptually identical to toJson/fromJson.
func WriteJSON(w io.Writer, s Summary) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(summaryJSON{Total: s.Total, ByCategory: s.ByCategory})
}

// WriteTable prints a human-readable table. Already implemented — plain
// fmt formatting, nothing exercise-worthy here.
func WriteTable(w io.Writer, s Summary) error {
	categories := make([]string, 0, len(s.ByCategory))
	for c := range s.ByCategory {
		categories = append(categories, c)
	}
	sort.Strings(categories)

	for _, c := range categories {
		if _, err := fmt.Fprintf(w, "%-20s %10.2f\n", c, s.ByCategory[c]); err != nil {
			return err
		}
	}
	_, err := fmt.Fprintf(w, "%-20s %10.2f\n", "TOTAL", s.Total)
	return err
}
