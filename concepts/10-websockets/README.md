# WebSockets: connection lifecycle, read/write pumps, the Hub pattern

## Dart concept
`web_socket_channel` gives a Flutter client a `Stream`/`Sink` pair for
one WebSocket connection — one of the closer Dart↔Go matches, since a
WS connection is naturally a duplex stream either side.

## Dart implementation
See `dart/chat_client.dart` (a commented sketch — this repo has no
Flutter app wired to this module yet; that's `projects/04`).

## Go equivalent
Split the same way Stage 7's Postgres module was split:
1. **`hub.go`** (the exercise, testable, no real socket needed) — `Hub`
   is the single owner of the connected-clients set, coordinating
   register/unregister/broadcast entirely through channels. No mutex,
   because only `Run`'s goroutine ever touches the map.
2. **`pumps.go`** (reference code, needs a live `*websocket.Conn`) —
   `ReadPump`/`WritePump`/`ServeWS`, the two-goroutines-per-connection
   pattern and ping/pong heartbeats. Exercised for real in
   `projects/03-chat-server`.

Library: `github.com/gorilla/websocket` — matches the roadmap doc's own
Hub-pattern code samples directly. This is the second non-stdlib
dependency in `concepts/go.mod`, after `pgx`.

## Go implementation
See `go/hub.go` + `go/hub_test.go`, and `go/pumps.go`.

## Important differences
- A Flutter client manages exactly **one** connection; a Go server may
  manage thousands and must never let one slow client block delivery to
  everyone else — there's no Dart equivalent to reason from here.
- The Hub's `broadcast` case sends to each client's buffered `Send`
  channel via a **non-blocking** inner `select` with a `default:`
  branch — a full buffer means that client is too slow, so it gets
  dropped instead of blocking the whole Hub. This one line is the
  entire backpressure mechanism, and it's the difference between a chat
  server that degrades gracefully and one bad client wedging everyone.
- Two goroutines per connection (`ReadPump`, `WritePump`) exist because
  reading and writing on the same connection can't happen concurrently
  anyway, and separating them means a slow write never blocks an
  incoming read.
- Without ping/pong heartbeats, a half-open connection (client's WiFi
  died mid-session) looks alive to the server forever — leaking a
  goroutine pair and a Hub map entry per dead client. Dart's `Stream`
  subscription has no equivalent silent-leak failure mode.

## Exercise (Level 6 — Build)
`Hub.Run` has a `// TODO` gap: implement the `select` over `Register`,
`Unregister`, and `Broadcast` — see the TODO comment in `hub.go` for the
exact three cases, including the non-blocking backpressure send. Remove
the `t.Skip(...)` calls in `hub_test.go` once implemented, and run under
`go test -race`.
