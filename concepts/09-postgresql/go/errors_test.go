package tasksdb

import (
	"database/sql"
	"errors"
	"testing"
)

// TestClassifyErrorMapsNoRows is the exercise: sql.ErrNoRows must become
// ErrTaskNotFound. Skipped until the TODO in errors.go is filled in.
func TestClassifyErrorMapsNoRows(t *testing.T) {
	t.Skip("TODO(exercise): implement ClassifyError, then remove this Skip")

	err := ClassifyError("42", sql.ErrNoRows)
	if !errors.Is(err, ErrTaskNotFound) {
		t.Errorf("ClassifyError(sql.ErrNoRows) = %v, want errors.Is match for ErrTaskNotFound", err)
	}
}

// TestClassifyErrorWrapsOtherErrors proves a non-ErrNoRows error is
// wrapped with context, not silently replaced. Skipped until the TODO
// in errors.go is filled in.
func TestClassifyErrorWrapsOtherErrors(t *testing.T) {
	t.Skip("TODO(exercise): implement ClassifyError, then remove this Skip")

	connErr := errors.New("connection refused")
	err := ClassifyError("42", connErr)

	if !errors.Is(err, connErr) {
		t.Errorf("ClassifyError(connErr) = %v, want it to wrap connErr (errors.Is match)", err)
	}
	if errors.Is(err, ErrTaskNotFound) {
		t.Errorf("ClassifyError(connErr) = %v, want it NOT to match ErrTaskNotFound", err)
	}
}

// TestClassifyErrorNilIsNil covers the no-error path.
// Skipped until the TODO in errors.go is filled in.
func TestClassifyErrorNilIsNil(t *testing.T) {
	t.Skip("TODO(exercise): implement ClassifyError, then remove this Skip")

	if err := ClassifyError("42", nil); err != nil {
		t.Errorf("ClassifyError(nil) = %v, want nil", err)
	}
}
