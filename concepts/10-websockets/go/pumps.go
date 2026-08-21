package chathub

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait  = 10 * time.Second
	pongWait   = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// WSClient pairs a Hub Client with a real WebSocket connection — the
// piece that needs a live connection to run, which is why this whole
// file is reference code rather than a unit-testable exercise (see
// hub.go for the testable part). It gets exercised for real in
// projects/03-chat-server. Embedding *Client (not copying its fields)
// matters: Hub tracks clients by pointer identity, so ReadPump's
// deferred Unregister must send the exact same *Client that was
// registered.
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
// the hub's broadcast channel. Two goroutines per connection (this one
// and WritePump) exist because reading and writing on the same
// underlying connection can't happen concurrently anyway, and
// separating them means a slow write never blocks an incoming read.
func (c *WSClient) ReadPump(hub *Hub) {
	defer func() {
		hub.Unregister <- c.Client
		c.Conn.Close()
	}()
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			break // client disconnected or protocol error
		}
		hub.Broadcast <- message
	}
}

// WritePump drains c.Send and writes to the socket, sending ping frames
// on an interval. Without this, a half-open connection (client's WiFi
// died mid-session) looks alive to the server forever, leaking a
// goroutine pair and a Hub map entry per dead client.
func (c *WSClient) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
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

// ServeWS upgrades an HTTP request to a WebSocket connection, registers
// the resulting client with hub, and starts its two pump goroutines.
func ServeWS(hub *Hub, w http.ResponseWriter, r *http.Request) error {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	client := NewWSClient(conn)
	hub.Register <- client.Client

	go client.WritePump()
	go client.ReadPump(hub)
	return nil
}
