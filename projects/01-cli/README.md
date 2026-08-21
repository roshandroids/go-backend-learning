# Project 1 — CLI Tool

## What I built
Scaffolded: CSV parsing (`csv.go`) and summary computation
(`summary.go`) are left as `// TODO` gaps — the exercise is filling
those in, not this wiring. `main.go` (flag parsing, file open, format
dispatch) and `output.go` (JSON/table formatting) are fully implemented
since routing/formatting are "skim" topics, same split used throughout
`concepts/`.

## Why I built it
Stage 1-3 fundamentals (packages, structs, methods, error handling) with
zero networking, per the roadmap.

## Architecture
Single `main` package -- no layers needed at this size.

## Technologies
Go stdlib only (`flag` or `os.Args`, `encoding/csv`, `encoding/json`).

## What I learned
_TBD — fill in once the exercise is done._

## Important design decisions
- `run(path, format string, out io.Writer)` takes an `io.Writer`, not
  `*os.File` — the accept-an-interface idiom (Stage 2), and what makes
  `main_test.go` able to assert on output without touching stdout.
- `ParseTransactions` fails the whole call on the first malformed row
  rather than skipping it — per the project's own completion criteria,
  a silent skip would hide bad data.

## Problems encountered
_TBD_

## How I solved them
_TBD_

## Exercise gaps (Level 2/6 — Complete/Build)
Two `// TODO`s:
1. `csv.go`'s `ParseTransactions` — parse each data row, wrapping a bad
   `amount` column into an error that names the row number (see the
   TODO comment for the exact shape `csv_test.go` expects).
2. `summary.go`'s `Summarize` — accumulate `Total` and `ByCategory`.

Remove the `t.Skip(...)` calls across `csv_test.go`, `summary_test.go`,
and `main_test.go` once both are implemented.

## How to run it
```
go run main.go <path-to-csv>
```

## Tests
```
go test -race ./...
```

## Future improvements
_TBD_
