package hub

import "sync"

// RoomManager owns one Hub per room ID, created lazily on first use —
// WS ladder rung 7 (rooms) is "multiple hubs, keyed by room ID," per
// the roadmap's own description of the pattern.
type RoomManager struct {
	mu    sync.Mutex
	rooms map[string]*Hub
}

func NewRoomManager() *RoomManager {
	return &RoomManager{rooms: make(map[string]*Hub)}
}

// GetOrCreate returns the Hub for roomID, creating and starting one
// (go hub.Run()) if this is the first time roomID has been seen.
//
// TODO(exercise, Level 4 — Refactor): lock rm.mu. If rm.rooms[roomID]
// already exists, return it. Otherwise create one with NewHub(), start
// it with `go h.Run()`, store it in rm.rooms[roomID], and return it.
// Unlock before returning (or use defer).
func (rm *RoomManager) GetOrCreate(roomID string) *Hub {
	// TODO
	return nil
}
