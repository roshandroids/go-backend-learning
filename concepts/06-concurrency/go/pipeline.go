// Package pipeline holds Stage 4's fan-out/fan-in exercise: multiple
// worker goroutines read from one input channel (fan-out), and their
// results land on a single output channel (fan-in) — the pattern behind
// almost every concurrent Go backend task.
package pipeline

import "context"

// Generate emits 0..n-1 on a buffered channel, then closes it. Buffered
// so the producer never blocks — even if a cancelled context means a
// downstream reader stops draining it early, this never leaks a
// goroutine waiting to send.
func Generate(n int) <-chan int {
	out := make(chan int, n)
	for i := 0; i < n; i++ {
		out <- i
	}
	close(out)
	return out
}

// Square fans out to numWorkers goroutines, each squaring values read
// from in, and fans their results back in on the returned channel.
//
// TODO(exercise, Level 6 — Build): implement the fan-out/fan-in pattern:
//  1. Launch numWorkers goroutines. Each one ranges over `in`, squares
//     every value, and sends it to `out`.
//  2. Use a sync.WaitGroup to track when all workers are done.
//  3. Launch one more goroutine that calls wg.Wait() then close(out) —
//     this is what lets a `for v := range out` on the caller's side
//     terminate instead of blocking forever.
//
// IMPORTANT: step 3's wg.Wait()+close(out) MUST run in its own goroutine,
// not inline in Square's body. Square itself must return immediately
// after launching the workers and the closer — otherwise Square blocks
// until all workers finish before ever returning `out`, and a caller
// that hasn't started reading `out` yet deadlocks against workers that
// are blocked trying to send into it.
//
// Run this under `go test -race` once implemented — a fan-out writing to
// a shared `out` channel is a classic spot to introduce (or fail to
// introduce) a data race.
func Square(in <-chan int, numWorkers int) <-chan int {
	out := make(chan int)
	// TODO: replace this placeholder close with the real fan-out/fan-in
	// implementation described above.
	close(out)
	return out
}

// SquareContext behaves like Square, but each worker also watches
// ctx.Done() and stops reading/squaring as soon as ctx is cancelled —
// the Go idiom for "stop working," in contrast to a Dart Stream
// subscription's implicit cancel.
//
// TODO(exercise, Level 6 — Build): same fan-out/fan-in shape as Square,
// but each worker's loop must select between receiving from `in` and
// ctx.Done(), returning immediately if ctx.Done() fires — even with
// unread values still sitting in `in`.
func SquareContext(ctx context.Context, in <-chan int, numWorkers int) <-chan int {
	out := make(chan int)
	// TODO: replace this placeholder close with the real fan-out/fan-in
	// implementation described above.
	close(out)
	return out
}
