package hub

import (
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

// These tests run a REAL WebSocket server (httptest.NewServer) and
// dial it with a real client, unlike hub_test.go's fake-Client unit
// tests. Backpressure/dropping a slow client is deliberately NOT
// exercised here over a real socket — reproducing it deterministically
// needs enough traffic to fill OS-level TCP buffers, which is exactly
// the kind of timing-dependent test that flakes in CI. That guarantee
// is already covered by TestHubDropsSlowClientInsteadOfBlocking's fake
// Client, which controls backpressure directly via an unbuffered
// channel instead of relying on socket buffering.

func newTestServer(t *testing.T) (*httptest.Server, *RoomManager) {
	t.Helper()
	rm := NewRoomManager()
	srv := httptest.NewServer(NewServer(rm))
	t.Cleanup(srv.Close)
	return srv, rm
}

func dialRoom(t *testing.T, srv *httptest.Server, room string) *websocket.Conn {
	t.Helper()
	url := "ws" + strings.TrimPrefix(srv.URL, "http") + "/ws?room=" + room
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		t.Fatalf("dial %s: %v", url, err)
	}
	t.Cleanup(func() { conn.Close() })
	return conn
}

// readUntil reads up to maxMessages from conn, within an overall
// timeout, looking for `want` — used because a presence "join" message
// may arrive before the message under test.
func readUntil(t *testing.T, conn *websocket.Conn, want string, maxMessages int) bool {
	t.Helper()
	conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	for i := 0; i < maxMessages; i++ {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return false
		}
		if string(msg) == want {
			return true
		}
	}
	return false
}

// TestSingleClientReceivesOwnBroadcast is the connect+echo milestone
// (WS ladder rungs 1-3): the Hub broadcasts to every registered client,
// including the sender. Skipped until Hub.Run and RoomManager.GetOrCreate
// are implemented.
func TestSingleClientReceivesOwnBroadcast(t *testing.T) {
	t.Skip("TODO(exercise): implement Hub.Run and RoomManager.GetOrCreate, then remove this Skip")

	srv, _ := newTestServer(t)
	conn := dialRoom(t, srv, "general")

	if err := conn.WriteMessage(websocket.TextMessage, []byte("hello")); err != nil {
		t.Fatalf("write: %v", err)
	}
	if !readUntil(t, conn, "hello", 3) {
		t.Error("did not receive own broadcast message back")
	}
}

// TestTwoClientsSameRoomSeeBroadcast is the multi-client broadcast
// milestone (WS ladder rungs 4-5). Skipped until Hub.Run and
// RoomManager.GetOrCreate are implemented.
func TestTwoClientsSameRoomSeeBroadcast(t *testing.T) {
	t.Skip("TODO(exercise): implement Hub.Run and RoomManager.GetOrCreate, then remove this Skip")

	srv, _ := newTestServer(t)
	a := dialRoom(t, srv, "general")
	b := dialRoom(t, srv, "general")

	if err := a.WriteMessage(websocket.TextMessage, []byte("hi from a")); err != nil {
		t.Fatalf("write: %v", err)
	}

	if !readUntil(t, a, "hi from a", 3) {
		t.Error("sender did not receive its own message")
	}
	if !readUntil(t, b, "hi from a", 3) {
		t.Error("other client in the same room did not receive the message")
	}
}

// TestClientsInDifferentRoomsAreIsolated is the rooms milestone (WS
// ladder rung 7): a message sent in one room must never reach a client
// in a different room. Skipped until Hub.Run and
// RoomManager.GetOrCreate are implemented.
func TestClientsInDifferentRoomsAreIsolated(t *testing.T) {
	t.Skip("TODO(exercise): implement Hub.Run and RoomManager.GetOrCreate, then remove this Skip")

	srv, _ := newTestServer(t)
	a := dialRoom(t, srv, "room-a")
	b := dialRoom(t, srv, "room-b")

	if err := a.WriteMessage(websocket.TextMessage, []byte("only for room-a")); err != nil {
		t.Fatalf("write: %v", err)
	}

	if !readUntil(t, a, "only for room-a", 3) {
		t.Error("sender in room-a did not receive its own message")
	}
	if readUntil(t, b, "only for room-a", 3) {
		t.Error("client in room-b received a message meant for room-a — rooms are not isolated")
	}
}
