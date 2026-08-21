package task

import "errors"

// ErrTitleRequired is returned when CreateRequest.Title is empty. This
// reuses the custom-error-value idiom from concepts/05-error-handling —
// no new concept, just applied here.
var ErrTitleRequired = errors.New("title is required")

// validateCreateRequest checks req before it reaches the repository.
//
// TODO(exercise, Level 2 — Complete): trim req.Title with
// strings.TrimSpace and return ErrTitleRequired if the result is empty.
// Return nil otherwise.
func validateCreateRequest(req CreateRequest) error {
	// TODO
	return nil
}
