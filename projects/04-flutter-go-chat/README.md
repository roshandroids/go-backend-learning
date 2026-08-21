# Project 4 — Flutter + Go Chat (the flagship project)

## What I built
Scaffolded Go-server-side only this session (no Flutter/live Postgres
available to actually run/test those parts here): `go_server/internal/chat`
has message-ID-and-ack (`ack.go`, `store.go`) and JWT validation
(`auth.go`) as `// TODO` gaps, both testable without a live database —
`auth_test.go` signs real JWTs with `golang-jwt/jwt/v5` to build its
fixtures. The Hub/room mechanics are deliberately NOT re-scaffolded a
third time here; `cmd/server/main.go` documents how to wire your own
already-implemented Hub from `projects/03-chat-server` together with
this package once both are ready. `flutter_app/` and a real Postgres
`MessageStore` remain future work.
## Why I built it
The thread connecting the whole curriculum: WebSocket challenge ladder
rungs 1-17, from a basic handshake through persistence, auth, and a
Redis-backed distributed system (Projects 7-9 in the roadmap).
## Architecture
`go_server/` -- extends Project 3's Hub with persistence + auth.
`flutter_app/` -- WS client using `web_socket_channel`.
Kept in the same repo, clearly separated, so client and backend are
always viewable together.
## Technologies
Go: gorilla/websocket, database/sql (Postgres), golang-jwt.
Flutter: web_socket_channel, uuid.
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
# terminal 1
cd go_server/cmd/server && go run main.go
# terminal 2
cd flutter_app && flutter run
```
## Tests
```
cd go_server && go test -race ./...
cd flutter_app && flutter test
```

## Exercise gaps (Level 2)
Two `// TODO`s in `go_server/internal/chat`:
1. `ack.go`'s `HandleIncoming` — persist via the store, return the ack.
2. `auth.go`'s `ValidateToken` — parse/validate an HS256 JWT, extract
   the `sub` claim, reject anything else as `ErrInvalidToken`.

Remove the `t.Skip(...)` calls in `ack_test.go` and `auth_test.go` once
both are implemented.

## Future improvements
_TBD_ — copy your Hub from `projects/03-chat-server` in and wire it with
this package (see `cmd/server/main.go`'s doc comment), build the Flutter
client, swap `InMemoryStore` for a real Postgres-backed one following
`concepts/09-postgresql`'s pattern, then Redis multi-instance work lands
here per challenges 15-17.
