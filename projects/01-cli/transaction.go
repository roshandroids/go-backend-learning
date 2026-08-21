package main

// Transaction is one row of the input CSV: date,category,amount.
type Transaction struct {
	Date     string
	Category string
	Amount   float64
}

// Summary is the computed output: overall total plus a per-category
// breakdown.
type Summary struct {
	Total      float64
	ByCategory map[string]float64
}
