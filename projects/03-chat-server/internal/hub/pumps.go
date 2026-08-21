package hub

import (
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait  = 10 * time.Second
	pongWait   = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
)

// WSClient pairs a Hub Client with a real WebSocket connection. Fully
// implemented — the read/write pump shape was already the exercise in
// concepts/10-websockets; here it's applied for real (WS ladder rungs
// 1-3, 9).
type WSClient struct {
	*Client
	Conn *websocket.Conn
}

func NewWSClient(conn *websocket.Conn) *WSClient {
	return &WSClient{
		Client: &Client{Send: make(chan []byte, 256)},
		Conn:   conn,
	}
}

// ReadPump reads from the client's socket and forwards each message to
// the hub's broadcast channel. Also arms the ping/pong heartbeat (WS
// ladder rung 9): without a pong within pongWait, the read deadline
// fires and ReadMessage returns an error, so a half-open connection
// (client's WiFi died) gets detected instead of leaking a goroutine
// pair forever.
func (c *WSClient) ReadPump(h *Hub) {
	defer func() {
		h.Unregister <- c.Client
		c.Conn.Close()
	}()

	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error {
		c.Conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			break // client disconnected, protocol error, or heartbeat timeout
		}
		h.Broadcast <- message
	}
}

// WritePump drains c.Send and writes to the socket, sending ping frames
// on an interval. Two goroutines per connection (this one and ReadPump)
// exist because reading and writing on the same connection can't happen
// concurrently anyway, and separating them means a slow write never
// blocks an incoming read.
func (c *WSClient) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
