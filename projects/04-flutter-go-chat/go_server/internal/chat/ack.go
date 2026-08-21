package chat

import "context"

// HandleIncoming persists msg via store and returns the ack to send
// back to the sender — WS ladder rung 13's message-ID-and-ack pattern.
//
// TODO(exercise, Level 2 — Complete): call store.Save(ctx, msg). If it
// errors, return a zero AckMessage and a wrapped error via
// fmt.Errorf("saving message %s: %w", msg.ID, err). On success, return
// AckMessage{Type: "ack", ID: msg.ID} and nil.
func HandleIncoming(ctx context.Context, store MessageStore, msg IncomingMessage) (AckMessage, error) {
	// TODO
	return AckMessage{}, nil
}
