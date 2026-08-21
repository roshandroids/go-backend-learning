package tasksdb

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	// Registers the "pgx" driver name with database/sql. This repo
	// teaches the portable database/sql API (same Learn/Compare/Test
	// loop works with any driver) while using pgx underneath — pgx is
	// the more actively maintained driver as of writing, versus the
	// older lib/pq.
	_ "github.com/jackc/pgx/v5/stdlib"
)

// NewPool opens a database/sql connection pool against Postgres and
// configures its limits. This is REFERENCE code — it compiles, but is
// not unit-tested here, because unit-testing a real connection pool
// needs a real Postgres (or testcontainers-go, deferred to Stage 9).
// It gets exercised for real once projects/04-flutter-go-chat wires up
// Postgres persistence.
//
// SetMaxOpenConns matters in a way sqflite never taught you to think
// about: your single-device, single-connection mobile database has no
// pool to exhaust. A backend under load absolutely does — Stage 7's
// exercise (see this module's README) has you deliberately set this to
// 2, fire 20 concurrent requests, and watch what happens.
func NewPool(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("opening pool: %w", err)
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)
	return db, nil
}

// TaskRepository queries tasks. Reference code — see NewPool's comment.
type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

// FindTitle returns the title of task id, using ClassifyError (the
// exercise above) to translate database/sql's sentinel error into
// ErrTaskNotFound.
func (r *TaskRepository) FindTitle(ctx context.Context, id string) (string, error) {
	var title string
	err := r.db.QueryRowContext(ctx,
		`SELECT title FROM tasks WHERE id = $1`, id,
	).Scan(&title)
	if err != nil {
		return "", ClassifyError(id, err)
	}
	return title, nil
}

// CreateWithNote demonstrates a transactional multi-table write — the
// pattern Stage 7's exercise asks you to build for real against a
// running Postgres instance.
func (r *TaskRepository) CreateWithNote(ctx context.Context, id, title, note string) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("beginning tx: %w", err)
	}
	defer tx.Rollback() // no-op once committed

	if _, err := tx.ExecContext(ctx,
		`INSERT INTO tasks (id, title) VALUES ($1, $2)`, id, title); err != nil {
		return fmt.Errorf("inserting task: %w", err)
	}
	if _, err := tx.ExecContext(ctx,
		`INSERT INTO task_notes (task_id, note) VALUES ($1, $2)`, id, note); err != nil {
		return fmt.Errorf("inserting note: %w", err)
	}
	return tx.Commit()
}
