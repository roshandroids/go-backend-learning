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
- [ ] Stage 1 — Go Fundamentals
  - [x] Point: struct, value receiver (`DistanceTo`), pointer receiver (`Translate`)
  - [ ] Point: `Scale` (value receiver) vs `ScaleInPlace` (pointer receiver)
- [ ] Stage 2 — Type System & Design Philosophy
- [ ] Stage 3 — Error Handling
- [ ] Stage 4 — Concurrency
- [ ] Stage 5 — HTTP / Backend Fundamentals
- [ ] Stage 6 — REST API Development
- [ ] Stage 7 — Databases
- [ ] Stage 8 — WebSockets
- [ ] Stage 9 — Testing
- [ ] Stage 10 — Production Backend Engineering
- [ ] Stage 11 — Architecture

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
