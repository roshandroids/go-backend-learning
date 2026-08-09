# Project 4 — Flutter + Go Chat (the flagship project)

## What I built
_TBD_
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
## Future improvements
_TBD_ (Redis multi-instance work lands here per challenges 15-17)
