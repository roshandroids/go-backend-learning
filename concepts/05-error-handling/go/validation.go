// Package validation holds Stage 3 exercises on custom error types and
// errors.As — the Go idiom for "catch this specific kind of failure"
// with no exception hierarchy involved.
package validation

import "fmt"

// ValidationError carries which field failed and why. This is the Go
// equivalent of a typed Dart exception (e.g. a custom
// ValidationException with a `field` getter), but there's no class
// hierarchy or `on X catch` — callers use errors.As to extract it.
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// ValidateAge returns a *ValidationError (as error) if age is out of a
// reasonable range, or nil.
//
// TODO(exercise, Level 2): return &ValidationError{Field: "age", Message:
// "must not be negative"} when age < 0, and &ValidationError{Field: "age",
// Message: "must be realistic"} when age > 150. Return nil otherwise.
func ValidateAge(age int) error {
	// TODO
	return nil
}

// DescribeError formats err for logging: if it's a *ValidationError,
// describe the specific field; otherwise fall back to a generic message.
// This mirrors Dart's `on ValidationException catch (e) { ... } catch
// (e) { ... }`, except there's no implicit stack unwinding to catch —
// errors.As is an explicit type assertion by value.
//
// TODO(exercise, Level 2): use errors.As to extract a *ValidationError
// from err. If it matches, return fmt.Sprintf("invalid %s: %s",
// ve.Field, ve.Message). Otherwise return fmt.Sprintf("unexpected error:
// %v", err).
func DescribeError(err error) string {
	// TODO
	return ""
}
