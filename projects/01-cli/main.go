// Command cli is Project 1: reads a CSV of transactions and prints a
// summary (total + per-category breakdown) as JSON or a table. Stdlib
// only, no CLI framework — flag alone is enough at this size.
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	format := flag.String("format", "table", "output format: table or json")
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "usage: cli [-format table|json] <path-to-csv>")
		os.Exit(2)
	}

	if err := run(flag.Arg(0), *format, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func run(path, format string, out io.Writer) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("opening %s: %w", path, err)
	}
	defer f.Close()

	txs, err := ParseTransactions(f)
	if err != nil {
		return fmt.Errorf("parsing %s: %w", path, err)
	}

	summary := Summarize(txs)

	switch format {
	case "json":
		return WriteJSON(out, summary)
	case "table":
		return WriteTable(out, summary)
	default:
		return fmt.Errorf("unknown format %q (want table or json)", format)
	}
}
