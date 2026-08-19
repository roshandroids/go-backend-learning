# Roadmap Progress

Full curriculum: see the companion document `../Backend_stack/docs/learning/go-learning-roadmap.md`.
Interactive teaching system: see `../Backend_stack/docs/learning/go-mentorship-protocol.md`.

> These live one level up, outside this repo and outside git (see brain page
> `external-curriculum-docs`). If you've cloned this repo elsewhere without that
> sibling folder, the links above will dangle — this repo's own README/ROADMAP
> and brain are self-sufficient for "what was built," just not for full stage
> sequencing/teaching methodology.

## You Are Here

- **Current stage:** Stage 1 — Go Fundamentals
- **Current exercise:** `concepts/01-fundamentals/README.md` — add `Scale`
  (value receiver, intentionally wrong) and `ScaleInPlace` (pointer receiver)
  to `Point`, plus a test proving the mutation difference.
- **Next command:** `cd concepts && go test ./... -run TestPoint -v`
- **Last touched:** 2026-08-10 (brain memory system added; last real code
  change was 2026-08-08).

## Stage Checklist

Each sub-item is the concrete first exercise queued for that concept folder
(all currently `_TBD_` stubs except 01). Filled in one at a time, in order,
following the mentorship protocol loop (Explain → Tiny Example → Predict →
Implement → Break → Debug → Compare → Apply). Execution starts when we
actually work through each one — this list is the plan, not a claim of
progress.

- [ ] Stage 1 — Go Fundamentals (`concepts/01-fundamentals`)
  - [x] Point: struct, value receiver (`DistanceTo`), pointer receiver (`Translate`)
  - [ ] Point: `Scale` (value receiver) vs `ScaleInPlace` (pointer receiver) — in progress
- [ ] Stage 2 — Type System & Design Philosophy
  - [ ] `concepts/02-types-and-structs`: zero values, nil-vs-empty slice/map footgun (`Store.Add` — scaffolded, TODO gap open)
  - [ ] `concepts/03-functions-and-methods`: first-class functions/closures (`Validator`/`All` — scaffolded, TODO gap open)
  - [ ] `concepts/04-interfaces`: consumer-defined `notifier` interface (`Alert` — scaffolded, TODO gap open) + embedding trap (`Animal`/`Dog` — implemented, Predict exercise, no gap)
- [ ] Stage 3 — Error Handling
  - [ ] `concepts/05-error-handling`: `ValidateAge`/`DescribeError`, custom `ValidationError` + `errors.As` (scaffolded, TODO gaps open)
- [ ] Stage 4 — Concurrency (major section — expect multiple passes)
  - [ ] `concepts/06-concurrency`: fan-out/fan-in pipeline (100 ints → 5 squarer workers → collector), clean under `-race`
  - [ ] `concepts/06-concurrency`: same pipeline, rewritten to respect `ctx.Done()` for clean cancellation
- [ ] Stage 5 — HTTP / Backend Fundamentals
  - [ ] `concepts/07-http`: stdlib-only server (`/health`, `/users/{id}`, `POST /users`), logging middleware, graceful `SIGTERM` shutdown
- [ ] Stage 6 — REST API Development
  - [ ] `concepts/08-rest-api-design`: extend Stage 5 server into in-memory CRUD `Task` API, validation, structured `APIError`, cursor pagination
- [ ] Stage 7 — Databases
  - [ ] `concepts/09-postgresql`: pooled repository, one transactional multi-table write, deliberate pool-exhaustion demo (`SetMaxOpenConns(2)`)
- [ ] Stage 8 — WebSockets
  - [ ] `concepts/10-websockets`: Hub/Client pattern in isolation (read pump/write pump/register/unregister) before it becomes `projects/03-chat-server`
- [ ] Stage 9 — Testing (cross-cutting, not a dedicated folder)
  - [ ] Applied retroactively per module as each is built: table-driven tests, `httptest`, eventually `testcontainers-go` for Stage 7
- [ ] Stage 10 — Production Backend Engineering (deferred, just-in-time)
  - [ ] Applied directly to `projects/04-flutter-go-chat` once it exists — not front-loaded, per roadmap's own guidance
- [ ] Stage 11 — Architecture
  - [ ] Applied as a retrospective ADR once a project outgrows the flat handler/service/repository layout — not scheduled yet

## Project Build Order

Scaffolded now: `01-cli`, `02-rest-api`, `03-chat-server`, `04-flutter-go-chat`.
Not yet scaffolded (roadmap Projects 5, 8, 9): Authenticated API (JWT + rate
limiting), production hardening (Docker/metrics/health checks), and the
Redis-backed distributed WS system — deliberately deferred, per the
roadmap's "don't front-load Stage 10/11" guidance. Add them when we're
actually about to need them, not before.

1. `projects/01-cli` — needs Stage 1/2/3 (structs, zero values, error handling) solid first.
2. `projects/02-rest-api` — needs Stage 5/6 (`concepts/07-http`, `08-rest-api-design`).
3. `projects/03-chat-server` — needs Stage 4/8 (`concepts/06-concurrency`, `10-websockets`); feeds WS Ladder items 1–12.
4. `projects/04-flutter-go-chat` — needs project 3 done + Stage 7 (`09-postgresql`) for persistence; feeds WS Ladder items 13–14.

## WebSocket Challenge Ladder Progress
- [ ] 1. Connect Flutter to Go
- [ ] 2. Send one message
- [ ] 3. Echo messages
- [ ] 4. Connect two clients
- [ ] 5. Broadcast messages
- [ ] 6. Implement the Hub
- [ ] 7. Add rooms
- [ ] 8. Add presence
- [ ] 9. Add ping/pong
- [ ] 10. Handle reconnects
- [ ] 11. Introduce a slow client
- [ ] 12. Implement backpressure
- [ ] 13. Persist messages in PostgreSQL
- [ ] 14. Add authentication
- [ ] 15. Run two Go server instances
- [ ] 16. Introduce Redis Pub/Sub
- [ ] 17. Test failure of one server
