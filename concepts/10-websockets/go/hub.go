// Package chathub holds Stage 8's Hub-pattern exercise: the central
// goroutine that owns every client connection and coordinates them
// entirely via channels — "don't communicate by sharing memory; share
// memory by communicating." Hub itself never touches a network
// connection, which is exactly what makes it testable without a real
// socket. The parts that DO need a real *websocket.Conn (readPump,
// writePump) live in pumps.go as reference code, exercised for real in
// projects/03-chat-server.
package chathub

// Client is anything the Hub can send bytes to via a buffered outbound
// queue. The real Client in projects/03-chat-server additionally wraps
// a *websocket.Conn — Hub only ever sees this shape.
type Client struct {
	Send chan []byte
}

// Hub is the single owner of the clients set. Because only Run's
// goroutine ever reads or writes `clients`, no mutex is needed —
// everything else talks to the Hub through these three channels.
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

// Run is the Hub's single event loop. Call it in its own goroutine.
//
// TODO(exercise, Level 6 — Build): implement the three cases of a
// `select` over h.Register, h.Unregister, and h.Broadcast, looping
// forever:
//  1. client := <-h.Register — add it to h.clients.
//  2. client := <-h.Unregister — if it's present in h.clients, delete it
//     and close(client.Send).
//  3. message := <-h.Broadcast — for every client in h.clients, attempt
//     `client.Send <- message` inside a NON-blocking inner select with a
//     `default:` branch. If the default fires, that client's buffer is
//     full (too slow): close(client.Send) and remove it from h.clients
//     instead of blocking delivery to everyone else. This one line is
//     the entire backpressure mechanism.
//
// Run this under `go test -race` once implemented.
func (h *Hub) Run() {
	// TODO
}
