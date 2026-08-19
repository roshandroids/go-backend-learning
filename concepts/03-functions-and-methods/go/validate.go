// Package validate holds Stage 2 exercises on first-class functions and
// closures, via a Validator function type — the same shape Stage 5's HTTP
// middleware will reuse.
package validate

import "errors"

// Task is the thing being validated.
type Task struct {
	Title string
}

// Validator checks a Task and returns an error if invalid, or nil.
// Functions are values in Go, same as in Dart — this is a named function
// type, not an interface with one method.
type Validator func(Task) error

// ErrEmptyTitle is returned when a Task's Title is empty.
var ErrEmptyTitle = errors.New("title must not be empty")

// RequireTitle is a Validator: any function with this exact signature
// satisfies the Validator type, no declaration needed at the definition site.
func RequireTitle(t Task) error {
	if t.Title == "" {
		return ErrEmptyTitle
	}
	return nil
}

// All composes multiple Validators into one. The returned function is a
// closure — it captures `validators` from All's enclosing scope and keeps
// referring to it every time it's called, long after All has returned.
//
// TODO(exercise, Level 2): run each validator against t, in order,
// returning the first non-nil error immediately (short-circuit). Return
// nil only if every validator passes.
func All(validators ...Validator) Validator {
	return func(t Task) error {
		// TODO: for _, v := range validators {
		//           if err := v(t); err != nil {
		//               return err
		//           }
		//       }
		//       return nil
		return nil
	}
}
