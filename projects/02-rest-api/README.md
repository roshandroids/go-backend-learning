# Project 2 — REST API

## What I built
_TBD_
## Why I built it
Stage 6 applied end-to-end: handler/service split, validation, structured
error responses, pagination -- in-memory first, Postgres added in Project 3.
## Architecture
`cmd/api` wires dependencies; `internal/task` holds handler/service/repository
for the one resource. Flat, not a Clean-Architecture-style layer split.
## Technologies
Go stdlib `net/http`, `encoding/json`.
## What I learned
_TBD_
## Important design decisions
_TBD_
## Problems encountered
_TBD_
## How I solved them
_TBD_
## How to run it
```
cd cmd/api && go run main.go
```
## Tests
```
go test -race ./...
```
## Future improvements
_TBD_
