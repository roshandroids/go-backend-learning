package chathub

import (
	"testing"
	"time"
)

const testTimeout = 2 * time.Second

func recvOrTimeout(t *testing.T, ch <-chan []byte) ([]byte, bool) {
	t.Helper()
	select {
	case v, ok := <-ch:
		return v, ok
	case <-time.After(testTimeout):
		t.Fatal("timed out waiting to receive — Hub.Run is likely stuck")
		return nil, false
	}
}

// TestHubBroadcastReachesAllClients is the exercise's happy path.
// Skipped until the TODO in hub.go is filled in.
func TestHubBroadcastReachesAllClients(t *testing.T) {
	t.Skip("TODO(exercise): implement Hub.Run, then remove this Skip")

	h := NewHub()
	go h.Run()

	a := &Client{Send: make(chan []byte, 1)}
	b := &Client{Send: make(chan []byte, 1)}
	h.Register <- a
	h.Register <- b

	h.Broadcast <- []byte("hello")

	if got, ok := recvOrTimeout(t, a.Send); !ok || string(got) != "hello" {
		t.Errorf("a.Send got %q, ok=%v, want %q, ok=true", got, ok, "hello")
	}
	if got, ok := recvOrTimeout(t, b.Send); !ok || string(got) != "hello" {
		t.Errorf("b.Send got %q, ok=%v, want %q, ok=true", got, ok, "hello")
	}
}

// TestHubUnregisterClosesSend proves a disconnected client's Send
// channel is closed so its writePump (in the real Client) can exit.
// Skipped until the TODO in hub.go is filled in.
func TestHubUnregisterClosesSend(t *testing.T) {
	t.Skip("TODO(exercise): implement Hub.Run, then remove this Skip")

	h := NewHub()
	go h.Run()

	c := &Client{Send: make(chan []byte, 1)}
	h.Register <- c
	h.Unregister <- c

	if _, ok := recvOrTimeout(t, c.Send); ok {
		t.Error("c.Send was not closed after Unregister")
	}
}

// TestHubDropsSlowClientInsteadOfBlocking is the exercise's core lesson:
// one client whose buffer is full must never block delivery to everyone
// else. Skipped until the TODO in hub.go is filled in.
func TestHubDropsSlowClientInsteadOfBlocking(t *testing.T) {
	t.Skip("TODO(exercise): implement Hub.Run, then remove this Skip")

	h := NewHub()
	go h.Run()

	slow := &Client{Send: make(chan []byte)} // unbuffered, nobody ever reads it
	fast := &Client{Send: make(chan []byte, 1)}
	h.Register <- slow
	h.Register <- fast

	// slow's buffer is immediately "full" (unbuffered, no reader), so
	// this broadcast must drop slow rather than blocking on it.
	h.Broadcast <- []byte("first")

	if got, ok := recvOrTimeout(t, fast.Send); !ok || string(got) != "first" {
		t.Fatalf("fast.Send got %q, ok=%v, want \"first\", ok=true — Hub blocked on the slow client", got, ok)
	}
	if _, ok := recvOrTimeout(t, slow.Send); ok {
		t.Error("slow.Send was not closed after being dropped for backpressure")
	}

	// Prove the Hub itself is still alive and not wedged: a second
	// broadcast must still reach the remaining client.
	h.Broadcast <- []byte("second")
	if got, ok := recvOrTimeout(t, fast.Send); !ok || string(got) != "second" {
		t.Errorf("fast.Send got %q, ok=%v, want %q, ok=true", got, ok, "second")
	}
}
