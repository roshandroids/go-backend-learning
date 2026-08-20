# Concurrency: goroutines, channels, select, context, sync primitives

## Dart concept
`Future.wait` runs a list of Futures concurrently on the same
single-threaded event loop; a `StreamSubscription.cancel()` stops
delivery to one listener.

## Dart implementation
See `dart/pipeline.dart`.

## Go equivalent
A fan-out/fan-in pipeline: `Generate` produces values on a channel,
`Square` fans out to several worker goroutines that each square values
and fan their results back in on one output channel, and `SquareContext`
adds `ctx.Done()` cancellation so workers stop early instead of running
to completion regardless.

## Go implementation
See `go/pipeline.go` and `go/pipeline_test.go`.

## Important differences
- Dart's `Future`s are cooperative and run on one thread per isolate; Go
  goroutines run on real OS threads in parallel and share memory by
  default — a fan-out writing to a shared channel/map without
  synchronization is a genuine data race, not just a style smell. Run
  `go test -race` once the exercise is implemented to confirm there isn't
  one.
- A channel is consumed by exactly one receiver — there's no `.listen()`
  with multiple independent subscribers unless you build that fan-out
  yourself (which is exactly what `Square` does, explicitly).
- `context.Context` cancellation propagates *forward* through an explicit
  call chain (as a parameter you thread through every function that might
  block), unlike a Dart subscription cancel, which only stops delivery to
  one listener and doesn't reach back into whatever produced the stream.
- An unbuffered channel send with no receiver blocks the sender forever —
  this is why `Generate` uses a *buffered* channel: so the producer can
  never leak a goroutine waiting to send into a pipeline nobody's draining
  anymore.

## Exercise (Level 6 — Build, major section)
Two `// TODO` gaps in `pipeline.go`:
1. `Square` — implement the fan-out/fan-in pattern (worker goroutines +
   `sync.WaitGroup` + a closer goroutine). Run `go test -race` once done.
2. `SquareContext` — same shape, but each worker must also select on
   `ctx.Done()` and exit early if it fires.

Remove the `t.Skip(...)` calls in `pipeline_test.go` once both are
implemented and passing under `-race`.
