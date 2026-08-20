# HTTP: net/http servers, middleware, request lifecycle

## Dart concept
You've only ever been an HTTP *client*. `dart/server.dart` sketches the
mirror image with `package:shelf` — Dart's closest equivalent to
`net/http` — to show middleware there is also just function composition,
via a commented sketch rather than a runnable file (shelf isn't a
dependency of this repo).

## Go equivalent
`net/http` handlers wired with stdlib `http.ServeMux` (Go 1.22+ path
patterns, no router library), and `LoggingMiddleware` — the concept that
deserves disproportionate attention here, since it has no Dio-interceptor
equivalent: middleware is a plain `func(http.Handler) http.Handler`.

## Go implementation
See `go/server.go` and `go/server_test.go`.

## Important differences
- There's no framework magic — middleware is a function that takes a
  `http.Handler` and returns one. Once this clicks, every third-party
  router/middleware library becomes instantly readable.
- The mental model closest to what you know isn't Dio — it's a Flutter
  `Widget build()`: a pure function from input (`Request`) to output
  (`Response`).
- JSON encode/decode and routing are "skim" topics (this repo's own
  `ROADMAP.md`/roadmap doc say so) — `HealthHandler`, `GetUserHandler`,
  and `CreateUserHandler` are already fully implemented; they're here to
  be read, not filled in.
- Graceful shutdown (`http.Server.Shutdown` + signal handling) has no
  Flutter equivalent — a Flutter app never needs to "drain" in-flight
  work before exiting. It's deliberately not exercised here as a unit
  test (signal handling doesn't fit `go test` well); it belongs in
  `projects/02-rest-api`'s `main.go` once that project is built out.

## Exercise (Level 2 — Complete)
`LoggingMiddleware` has a `// TODO` gap: record `time.Now()`, call
`next.ServeHTTP(w, r)`, then log method/path/duration via
`slog.Info(...)`. Remove the `t.Skip(...)` in
`TestLoggingMiddlewarePassesThroughToNext` once implemented.
