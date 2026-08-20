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
func TestSquareFanOutFanIn(t *testing.T) {
	t.Skip("TODO(exercise): implement Square, then remove this Skip")

	const n = 100
	in := Generate(n)
	out := Square(in, 5)

	sum := 0
	count := 0
	for v := range out {
		sum += v
		count++
	}

	if count != n {
		t.Fatalf("got %d results, want %d", count, n)
	}

	want := 0
	for i := 0; i < n; i++ {
		want += i * i
	}
	if sum != want {
		t.Errorf("sum of squares = %d, want %d", sum, want)
	}
}

// TestSquareContextStopsOnCancel proves cancelling ctx lets the returned
// channel close instead of hanging forever, even with unread work left
// in `in`. Skipped until the TODO in pipeline.go is filled in.
func TestSquareContextStopsOnCancel(t *testing.T) {
	t.Skip("TODO(exercise): implement SquareContext, then remove this Skip")

	ctx, cancel := context.WithCancel(context.Background())
	in := Generate(1000) // far more than we'll actually read
	out := SquareContext(ctx, in, 5)

	<-out
	<-out
	cancel()

	done := make(chan struct{})
	go func() {
		for range out {
		}
		close(done)
	}()

	select {
	case <-done:
		// out closed after cancellation — success.
	case <-time.After(2 * time.Second):
		t.Fatal("out channel never closed after context cancellation — likely a goroutine leak")
	}
}
