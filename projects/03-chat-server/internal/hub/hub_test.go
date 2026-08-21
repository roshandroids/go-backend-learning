package hub

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

// TestHubBroadcastReachesAllClients uses a fake Client (no real socket)
// to exercise Hub.Run's core concurrency logic quickly, same as
// concepts/10-websockets. Skipped until the TODO in hub.go is filled in.
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

// TestHubDropsSlowClientInsteadOfBlocking is the backpressure exercise,
// same as concepts/10-websockets. Skipped until the TODO in hub.go is
// filled in.
func TestHubDropsSlowClientInsteadOfBlocking(t *testing.T) {
	t.Skip("TODO(exercise): implement Hub.Run, then remove this Skip")

	h := NewHub()
	go h.Run()

	slow := &Client{Send: make(chan []byte)} // unbuffered, nobody ever reads it
	fast := &Client{Send: make(chan []byte, 1)}
	h.Register <- slow
	h.Register <- fast

	h.Broadcast <- []byte("first")

	if got, ok := recvOrTimeout(t, fast.Send); !ok || string(got) != "first" {
		t.Fatalf("fast.Send got %q, ok=%v, want \"first\", ok=true — Hub blocked on the slow client", got, ok)
	}
	if _, ok := recvOrTimeout(t, slow.Send); ok {
		t.Error("slow.Send was not closed after being dropped for backpressure")
	}
}

// TestRoomManagerReusesExistingHub is the exercise for GetOrCreate.
// Skipped until the TODO in roommanager.go is filled in.
func TestRoomManagerReusesExistingHub(t *testing.T) {
	t.Skip("TODO(exercise): implement RoomManager.GetOrCreate, then remove this Skip")

	rm := NewRoomManager()
	a := rm.GetOrCreate("general")
	b := rm.GetOrCreate("general")
	c := rm.GetOrCreate("random")

	if a != b {
		t.Error("GetOrCreate(\"general\") returned different Hubs on repeated calls")
	}
	if a == c {
		t.Error("GetOrCreate(\"random\") returned the same Hub as a different room")
	}
}
