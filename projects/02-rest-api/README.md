# Project 2 — REST API

## What I built
Scaffolded: `internal/task` has the flat handler/service/repository/model
split wired end-to-end and running (`POST /tasks`, `GET /tasks/{id}`,
`GET /tasks`). `validateCreateRequest` and `Repository.Create/Get/List`
are left as `// TODO` gaps — handler/routing/JSON (`handler.go`) and the
validate-then-delegate wiring (`service.go`) are fully implemented, same
split used throughout `concepts/`. Manually verified: a missing task
already correctly 404s; create/list currently return placeholder empty
values until the repository TODOs are filled in.
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

## Exercise gaps (Level 2/4/6)
Three `// TODO`s in `internal/task`:
1. `validation.go`'s `validateCreateRequest` — trim and require `Title`.
2. `repository.go`'s `Create`/`Get` — basic map CRUD.
3. `repository.go`'s `List` — cursor pagination, same algorithm as
   `concepts/08-rest-api-design`'s `Store.List`, applied here.

Remove the `t.Skip(...)` calls across `validation_test.go`,
`repository_test.go`, and `handler_test.go` once implemented.

## Future improvements
_TBD_ — authentication/rate limiting (Project 5) and Postgres
persistence (Project 4) are deliberately out of scope here.
