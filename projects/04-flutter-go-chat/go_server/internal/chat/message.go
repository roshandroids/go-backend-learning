// Package chat holds Project 4's genuinely new server-side material:
// message-ID-and-ack persistence (WS ladder rung 13) and JWT auth
// (rung 14). The Hub/room mechanics themselves are NOT re-scaffolded a
// third time here — copy your own working Hub/RoomManager/pumps from
// projects/03-chat-server's internal/hub once you've implemented them
// there, and wire this package's HandleIncoming/ValidateToken into that
// server's message flow and upgrade handshake respectively.
package chat

// IncomingMessage is what a client sends over the WebSocket: text plus
// a client-generated ID, so the server can ack the exact message that
// was persisted.
type IncomingMessage struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

// AckMessage is echoed back to the sender once IncomingMessage is
// durably stored.
type AckMessage struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}
