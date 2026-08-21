package main

// Summarize computes the overall total and per-category breakdown for
// txs.
//
// TODO(exercise, Level 2 — Complete): sum every transaction's Amount
// into Total, and accumulate per-Category sums into ByCategory (make
// the map before writing into it — a nil map panics on write, same
// footgun as concepts/02-types-and-structs).
func Summarize(txs []Transaction) Summary {
	s := Summary{ByCategory: make(map[string]float64)}
	// TODO: for _, t := range txs {
	//           s.Total += t.Amount
	//           s.ByCategory[t.Category] += t.Amount
	//       }
	return s
}
