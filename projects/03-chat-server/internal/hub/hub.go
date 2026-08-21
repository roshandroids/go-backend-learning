// Package hub is Project 3: the Hub pattern from concepts/10-websockets,
// applied for real over a live WebSocket server this time, plus rooms
// (WS ladder rung 7), presence (rung 8), and ping/pong (rung 9).
package hub

// Client is anything the Hub can send bytes to via a buffered outbound
// queue. WSClient (in pumps.go) additionally wraps a real
// *websocket.Conn — Hub only ever sees this shape.
type Client struct {
	Send chan []byte
}

// Hub is the single owner of one room's connected-clients set — one Hub
// instance per room, managed by RoomManager (see roommanager.go).
// Because only Run's goroutine ever touches `clients`, no mutex is
// needed.
type Hub struct {
	clients    map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

// NewHub returns a ready-to-run Hub. Call go hub.Run() before using it.
func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

// Run is the Hub's single event loop — same shape as
// concepts/10-websockets' exercise, re-implemented here against a real
// server (WS ladder rung 6).
//
// TODO(exercise, Level 6 — Build): implement the three cases of a
// `select` over h.Register, h.Unregister, and h.Broadcast, looping
// forever:
//  1. client := <-h.Register — add it to h.clients.
//  2. client := <-h.Unregister — if present, delete it and
//     close(client.Send).
//  3. message := <-h.Broadcast — for every client in h.clients, attempt
//     `client.Send <- message` inside a non-blocking inner select with a
//     `default:` branch. If default fires, that client is too slow:
//     close(client.Send) and remove it from h.clients instead of
//     blocking delivery to everyone else.
//
// Run under `go test -race` once implemented.
func (h *Hub) Run() {
	// TODO
}
