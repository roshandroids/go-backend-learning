// Command server is Project 3: a standalone WebSocket broadcast server
// with rooms, presence, and ping/pong heartbeats — the Hub pattern from
// Stage 8, applied for real over a live server this time.
package main

import (
	"log"
	"net/http"

	"github.com/YOUR_USERNAME/go-backend-learning/projects/03-chat-server/internal/hub"
)

func main() {
	rm := hub.NewRoomManager()
	mux := hub.NewServer(rm)

	log.Println("listening on :8081")
	log.Fatal(http.ListenAndServe(":8081", mux))
}
