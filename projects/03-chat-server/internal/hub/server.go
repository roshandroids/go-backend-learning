package hub

import (
	"encoding/json"
	"log"
	"net/http"
	"sync/atomic"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true }, // learning project, not production
}

var nextClientID int64

// presenceEvent is broadcast to a room whenever a client joins or
// leaves — WS ladder rung 8, derived directly from register/unregister
// rather than tracked separately.
type presenceEvent struct {
	Type     string `json:"type"` // "join" or "leave"
	ClientID int64  `json:"client_id"`
}

// ServeWS upgrades an HTTP request to a WebSocket connection, registers
// the resulting client with the room's Hub (creating the room on first
// use — WS ladder rung 7), announces presence, and starts its two pump
// goroutines. Room is chosen by the `room` query parameter
// (?room=general), defaulting to "general".
func ServeWS(rm *RoomManager, w http.ResponseWriter, r *http.Request) {
	room := r.URL.Query().Get("room")
	if room == "" {
		room = "general"
	}
	h := rm.GetOrCreate(room)

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}

	client := NewWSClient(conn)
	clientID := atomic.AddInt64(&nextClientID, 1)
	h.Register <- client.Client

	if joinMsg, err := json.Marshal(presenceEvent{Type: "join", ClientID: clientID}); err == nil {
		h.Broadcast <- joinMsg
	}

	go client.WritePump()
	go func() {
		client.ReadPump(h)
		if leaveMsg, err := json.Marshal(presenceEvent{Type: "leave", ClientID: clientID}); err == nil {
			h.Broadcast <- leaveMsg
		}
	}()
}

// NewServer wires the /ws upgrade endpoint against rm.
func NewServer(rm *RoomManager) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ServeWS(rm, w, r)
	})
	return mux
}
