package chat

import (
	"context"
	"sync"
)

// MessageStore is defined at the CONSUMER (this package), narrow —
// Stage 2's idiom, applied a third time across this repo. InMemoryStore
// below satisfies it for tests; a real Postgres-backed implementation
// is future work — follow concepts/09-postgresql's split (a testable
// error-classification piece plus reference-only pooled/transactional
// code) rather than duplicating it here, to keep this project's scope
// on the genuinely new material: ack persistence and JWT auth.
type MessageStore interface {
	Save(ctx context.Context, msg IncomingMessage) error
}

// InMemoryStore is a MessageStore for tests and local runs before a
// real database is wired in.
type InMemoryStore struct {
	mu       sync.Mutex
	messages []IncomingMessage
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{}
}

func (s *InMemoryStore) Save(ctx context.Context, msg IncomingMessage) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.messages = append(s.messages, msg)
	return nil
}

// All returns every saved message, for tests to assert against.
func (s *InMemoryStore) All() []IncomingMessage {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]IncomingMessage, len(s.messages))
	copy(out, s.messages)
	return out
}
