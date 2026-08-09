# go-backend-learning

A working repository documenting a Flutter/Dart engineer's progression into production Go backend development — real code, real tests, real commits, not tutorial copies.

Curriculum lives in two companion documents (kept alongside this repo, linked here for reference):
- `ROADMAP.md` — stage-by-stage curriculum index and progress checkboxes
- `dart-vs-go.md` — the living Dart ↔ Go cross-reference table
- `docs/journal/` — one entry per real learning session
- `docs/decisions/` — short architecture decision records for the few calls worth writing down

## Current Capabilities
- [ ] Go syntax & program structure
- [ ] Structs & methods
- [ ] Interfaces (consumer-defined)
- [ ] Error handling
- [ ] Concurrency (goroutines/channels)
- [ ] HTTP servers
- [ ] REST API design
- [ ] PostgreSQL
- [ ] WebSockets
- [ ] Authentication
- [ ] Redis / distributed systems

## Projects Completed
- [ ] 01 — Go CLI
- [ ] 02 — REST API
- [ ] 03 — Chat server (Go only)
- [ ] 04 — Flutter + Go chat

## Repository Layout
```
concepts/     — one shared Go module; small, focused exercises per topic, paired with Dart where meaningful
projects/     — real, independently runnable applications, each with its own go.mod
docs/         — journal + architecture decision records
```

## Versioning
Tags mark curriculum milestones: v0.1 fundamentals → v0.2 REST API → v0.3 PostgreSQL →
v0.4 WebSocket server → v0.5 Flutter+Go chat → v0.6 auth → v0.7 Redis → v1.0 distributed chat.
