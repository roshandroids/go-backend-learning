// Package notify holds Stage 2 exercises on implicit interface
// satisfaction, consumer-defined interfaces, and the embedding-has-no-
// dynamic-dispatch trap.
package notify

// notifier is defined at the CONSUMER (this package), unexported and
// narrow — only the one method Alert actually needs. Neither
// EmailNotifier nor SMSNotifier ever declares "implements notifier"
// anywhere; Go checks this structurally, at compile time, wherever a
// notifier is required.
type notifier interface {
	Notify(message string) error
}

// EmailNotifier and SMSNotifier are concrete types with no interface in
// sight at their own definition — the opposite of Dart's idiom of
// defining an abstract class first (see dart/notify.dart).
type EmailNotifier struct {
	Address string
}

func (e EmailNotifier) Notify(message string) error {
	return nil // pretend to send an email
}

type SMSNotifier struct {
	Phone string
}

func (s SMSNotifier) Notify(message string) error {
	return nil // pretend to send an SMS
}

// Alert sends message via every notifier.
//
// TODO(exercise, Level 2): call Notify(message) on each notifier, in
// order, returning the first non-nil error immediately. Return nil only
// if every notifier succeeds.
func Alert(message string, notifiers ...notifier) error {
	// TODO: for _, n := range notifiers {
	//           if err := n.Notify(message); err != nil {
	//               return err
	//           }
	//       }
	//       return nil
	return nil
}
