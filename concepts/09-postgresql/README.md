# PostgreSQL: database/sql, pooling, transactions, migrations

## Dart concept
Closest analogue is Drift (compile-time-checked SQL, typed queries) more
than sqflite's raw string queries — but neither gives you a stdlib
sentinel error for "no rows" the way `sql.ErrNoRows` does.

## Dart implementation
See `dart/notfound.dart`.

## Go equivalent
Two things, deliberately split by whether they need a live database:
1. **`ClassifyError`** (testable, no DB needed) — translates
   `sql.ErrNoRows` into the domain-level `ErrTaskNotFound`, and wraps any
   other error with context. This is the actual exercise.
2. **`NewPool`/`TaskRepository`** (reference code, `repository.go`) — a
   real `database/sql` pool backed by the `pgx` driver, connection-limit
   configuration, and a transactional multi-table write. It compiles but
   is **not unit-tested here** — you cannot construct a working
   `*sql.Row` without a real Postgres connection (or a mocking library
   this repo doesn't depend on), so this gets exercised for real once
   `projects/04-flutter-go-chat` wires up Postgres persistence.

## Go implementation
See `go/errors.go` + `go/errors_test.go` (the exercise), and
`go/repository.go` (reference).

## Important differences
- `database/sql` is lower-level than Drift or sqflite — closer to
  writing raw SQL strings with `sqflite`, except pooling and prepared-
  statement caching are built into the standard library from day one.
- `SetMaxOpenConns` matters in a way sqflite never taught you to think
  about: a single-device, single-connection mobile database has no pool
  to exhaust. A backend under load absolutely does.
- `defer rows.Close()` (on `Query`/`QueryContext`, not shown here since
  this module uses `QueryRowContext`) has no sqflite equivalent — Go
  leaks connections if you forget it; sqflite manages cursor lifecycle
  for you.
- The Go community favors raw `database/sql` + a thin mapping layer
  (`sqlc`) over a full ORM (`gorm`) — don't reach for an ORM as your
  first tool; it hides exactly the pooling/N+1-query mechanics you need
  to understand.

## Exercise (Level 2 — Complete)
`ClassifyError` has a `// TODO` gap: if `errors.Is(err, sql.ErrNoRows)`,
return `ErrTaskNotFound`; otherwise wrap `err` with
`fmt.Errorf("finding task %s: %w", id, err)`; return `nil` for a `nil`
input. Remove the `t.Skip(...)` calls in `errors_test.go` once
implemented.

**When you have a real Postgres running** (not required for this
module): point `NewPool` at it, run `TaskRepository.FindTitle`/
`CreateWithNote` against a real `tasks`/`task_notes` schema, and
deliberately set `SetMaxOpenConns(2)` while firing 20 concurrent
requests to observe pool exhaustion — the Stage 7 exercise from the
external roadmap doc, intentionally deferred to `projects/04` where
Postgres is actually reachable.
