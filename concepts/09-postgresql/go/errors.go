// Package tasksdb holds Stage 7 exercises. The testable exercise below
// is deliberately decoupled from a live database connection — you
// cannot construct a working *sql.Row outside package database/sql, so
// unit-testing code that calls QueryRowContext directly requires either
// a real Postgres or a mocking library. Neither is worth adding for a
// concept module; the real, connected version lives in repository.go
// (see its comment) and gets exercised for real in projects/04 once
// Postgres is actually reachable there.
package tasksdb

import "errors"

// ErrTaskNotFound is the domain-level error the service layer checks for
// — it never needs to import database/sql to know "not found" happened.
var ErrTaskNotFound = errors.New("task not found")

// ClassifyError translates a database/sql error into a domain error:
// sql.ErrNoRows becomes ErrTaskNotFound; anything else is wrapped with
// context via fmt.Errorf's %w, preserving errors.Is/As compatibility —
// the same idiom concepts/05-error-handling introduced, applied here to
// the specific sentinel error database/sql uses for "no rows".
//
// TODO(exercise, Level 2): if errors.Is(err, sql.ErrNoRows), return
// ErrTaskNotFound. Otherwise return fmt.Errorf("finding task %s: %w",
// id, err). Return nil if err is nil.
func ClassifyError(id string, err error) error {
	// TODO
	return nil
}
