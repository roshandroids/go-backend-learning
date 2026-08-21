# Project 3 — Chat Server (Go only, no Flutter yet)

## What I built
Scaffolded: `internal/hub` has `Hub` (register/unregister/broadcast +
backpressure), `RoomManager` (one Hub per room, lazy-created),
`WSClient`'s read/write pumps (ping/pong heartbeat, fully implemented),
and `server.go` wiring `/ws?room=X` with presence join/leave broadcasts.
`Hub.Run` and `RoomManager.GetOrCreate` are `// TODO` gaps — everything
else is fully implemented. Build verified (`go build ./cmd/server`);
the real WS integration tests are written but skipped until the gaps
are filled in, deliberately not run against a live client manually
this session (unlike Projects 1-2) since that would require exercising
the same unimplemented Hub logic the tests already cover once unskipped.
## Why I built it
WebSocket challenge ladder rungs 1-9: connect, echo, broadcast, the Hub
pattern, rooms, presence, ping/pong -- proven out standalone before
Flutter integration (Project 4). Rungs 11-12 (slow client, backpressure)
are covered by the fake-Client unit tests in `hub_test.go`, not the real
socket integration tests, deliberately -- reproducing backpressure
deterministically over a real OS socket is exactly the kind of
timing-dependent test that flakes in CI.
## Architecture
`internal/hub` owns all client connection state; `cmd/server` wires the
HTTP upgrade handler and starts the Hub's run loop.
## Technologies
`github.com/gorilla/websocket`, Go stdlib.
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
cd cmd/server && go run main.go
```
## Tests
```
go test -race ./...
```

## Exercise gaps (Level 4/6)
Two `// TODO`s in `internal/hub`:
1. `hub.go`'s `Hub.Run` — same shape as `concepts/10-websockets`,
   re-implemented here against a real server.
2. `roommanager.go`'s `RoomManager.GetOrCreate` — lazily create and
   start a Hub per room ID.

Remove the `t.Skip(...)` calls across `hub_test.go` and
`integration_test.go` once both are implemented, and run under
`go test -race`.

## Future improvements
_TBD_ — reconnect handling (WS ladder rung 10) and multi-instance/Redis
(rungs 15-17) are deliberately out of scope here; see `projects/04` and
beyond.
