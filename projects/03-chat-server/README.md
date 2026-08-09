# Project 3 — Chat Server (Go only, no Flutter yet)

## What I built
_TBD_
## Why I built it
WebSocket challenge ladder rungs 1-12: Hub pattern, rooms, backpressure,
heartbeats -- proven out standalone before Flutter integration (Project 4).
## Architecture
`internal/hub` owns all client connection state; `cmd/server` wires the
HTTP upgrade handler and starts the Hub's run loop.
## Technologies
`github.com/gorilla/websocket` (or `nhooyr.io/websocket`), Go stdlib.
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
## Future improvements
_TBD_
