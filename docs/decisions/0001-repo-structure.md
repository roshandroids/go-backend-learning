# ADR 0001: Concepts module vs per-project modules

## Decision
`concepts/` uses one shared `go.mod` for all learning-snippet folders.
Each entry under `projects/` gets its own `go.mod`.

## Why
Concept folders are throwaway teaching artifacts that will never be
deployed -- one module keeps `go test ./...` trivial across all of them.
Projects are real, independently runnable things (some later
containerized) -- separate modules mirror how real services are
structured and prevent one project's dependency bump from touching
another's `go.sum`.

## Alternatives considered
- One repo-wide module: rejected -- would force every concept folder to
  share every project's dependencies (Postgres driver, JWT lib, etc.)
  even when irrelevant to that concept.
- One module per concept folder (10+ modules): rejected -- unnecessary
  overhead for snippets this small.
