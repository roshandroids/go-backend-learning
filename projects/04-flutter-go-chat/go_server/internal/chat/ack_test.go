package chat

import (
	"context"
	"errors"
	"strings"
	"testing"
)

type fakeStore struct {
	err error
}

func (s fakeStore) Save(ctx context.Context, msg IncomingMessage) error { return s.err }

// TestHandleIncomingSavesAndAcks is the exercise's happy path. Skipped
// until the TODO in ack.go is filled in.
func TestHandleIncomingSavesAndAcks(t *testing.T) {
	t.Skip("TODO(exercise): implement HandleIncoming, then remove this Skip")

	store := NewInMemoryStore()
	msg := IncomingMessage{ID: "1", Text: "hello"}

	ack, err := HandleIncoming(context.Background(), store, msg)
	if err != nil {
		t.Fatalf("HandleIncoming() error = %v, want nil", err)
	}
	if ack.Type != "ack" || ack.ID != "1" {
		t.Errorf("ack = %+v, want {ack 1}", ack)
	}

	saved := store.All()
	if len(saved) != 1 || saved[0] != msg {
		t.Errorf("store.All() = %+v, want [%+v]", saved, msg)
	}
}

// TestHandleIncomingWrapsStoreError proves a storage failure surfaces
// with the message ID in context, not a bare error. Skipped until the
// TODO in ack.go is filled in.
func TestHandleIncomingWrapsStoreError(t *testing.T) {
	t.Skip("TODO(exercise): implement HandleIncoming, then remove this Skip")

	saveErr := errors.New("connection refused")
	store := fakeStore{err: saveErr}
	msg := IncomingMessage{ID: "42", Text: "hello"}

	_, err := HandleIncoming(context.Background(), store, msg)
	if err == nil {
		t.Fatal("HandleIncoming() error = nil, want a wrapped error")
	}
	if !errors.Is(err, saveErr) {
		t.Errorf("error = %v, want it to wrap the store error (errors.Is)", err)
	}
	if !strings.Contains(err.Error(), "42") {
		t.Errorf("error = %q, want it to mention the message ID (42)", err.Error())
	}
}
