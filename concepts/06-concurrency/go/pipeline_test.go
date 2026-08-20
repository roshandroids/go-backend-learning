package pipeline

import (
	"context"
	"testing"
	"time"
)

func TestGenerate(t *testing.T) {
	out := Generate(5)

	sum := 0
	count := 0
	for v := range out {
		sum += v
		count++
	}

	if count != 5 {
		t.Fatalf("got %d values, want 5", count)
	}
	if sum != 0+1+2+3+4 {
		t.Errorf("sum = %d, want 10", sum)
	}
}

// TestSquareFanOutFanIn is the exercise: every value from Generate must
// come back squared exactly once, regardless of which of the 5 workers
// happened to process it. Skipped until the TODO in pipeline.go is
// filled in. Run with `go test -race` once it's not skipped.
//
// The collection loop runs in its own goroutine with a timeout guard: a
// common bug here is closing `out` inline (`wg.Wait(); close(out)`)
// instead of in its own goroutine, which makes Square itself never
// return. Without this guard that bug hangs the test forever with no
// diagnostic; with it, you get a clear failure message instead.
func TestSquareFanOutFanIn(t *testing.T) {
	t.Skip("TODO(exercise): implement Square, then remove this Skip")

	const n = 100
	in := Generate(n)
	out := Square(in, 5)

	type result struct {
		sum, count int
	}
	done := make(chan result, 1)
	go func() {
		sum, count := 0, 0
		for v := range out {
			sum += v
			count++
		}
		done <- result{sum, count}
	}()

	select {
	case r := <-done:
		if r.count != n {
			t.Fatalf("got %d results, want %d", r.count, n)
		}
		want := 0
		for i := 0; i < n; i++ {
			want += i * i
		}
		if r.sum != want {
			t.Errorf("sum of squares = %d, want %d", r.sum, want)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("out channel never closed — Square likely blocks instead of returning " +
			"(close(out) must happen in its own goroutine after wg.Wait(), not inline)")
	}
}

// TestSquareContextStopsOnCancel proves cancelling ctx actually truncates
// the work — not just that the channel eventually closes, which a
// correct-but-uncancellable implementation would also satisfy. Skipped
// until the TODO in pipeline.go is filled in.
func TestSquareContextStopsOnCancel(t *testing.T) {
	t.Skip("TODO(exercise): implement SquareContext, then remove this Skip")

	const n = 1000
	ctx, cancel := context.WithCancel(context.Background())
	in := Generate(n)
	out := SquareContext(ctx, in, 5)

	count := 0
	for i := 0; i < 2; i++ {
		v, ok := <-out
		if !ok {
			t.Fatalf("out closed after only %d values, want at least 2", i)
		}
		if isqrt := isPerfectSquare(v); !isqrt {
			t.Errorf("got %d, want a perfect square (workers should square their input)", v)
		}
		count++
	}
	cancel()

	done := make(chan struct{})
	go func() {
		for range out {
			count++
		}
		close(done)
	}()

	select {
	case <-done:
		// out closed after cancellation.
	case <-time.After(2 * time.Second):
		t.Fatal("out channel never closed after context cancellation — likely a goroutine leak")
	}

	if count >= n {
		t.Errorf("consumed all %d values despite cancelling after 2 — cancellation didn't truncate the work", n)
	}
}

func isPerfectSquare(v int) bool {
	if v < 0 {
		return false
	}
	for k := 0; k*k <= v; k++ {
		if k*k == v {
			return true
		}
	}
	return false
}
