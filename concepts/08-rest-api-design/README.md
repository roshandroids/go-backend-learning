# REST API Design: handler/service/repository, validation, pagination

## Dart concept
A `Store.list(cursor, limit)` that returns the next N tasks after a given
cursor — the same idea as a Flutter `ListView` requesting "give me the
next N items after the last one I rendered," never "give me items 40-60."

## Dart implementation
See `dart/pagination.dart` (fully implemented — the Go version is the
exercise, not this one).

## Go equivalent
`Store.List(cursor, limit)` — cursor-based pagination over an in-memory
`Task` slice. Validation and structured error responses (`APIError`) are
deliberately **not** re-taught in this module; `concepts/05-error-handling`
already covers the custom-error-type idiom they'd repeat. This module
isolates the one genuinely new idea for Stage 6: pagination strategy.

## Go implementation
See `go/pagination.go` and `go/pagination_test.go`.

## Important differences
- Cursor-based pagination stays correct even if tasks are added/removed
  between page requests — offset pagination (`LIMIT`/`OFFSET`) doesn't;
  a row inserted before your current offset shifts every subsequent page.
- The reason this matters beyond correctness: once Stage 7 introduces a
  real database, `OFFSET` degrades badly on large tables because the
  database still has to scan and discard every skipped row — a cursor
  (typically an indexed column like `id` or `created_at`) avoids that scan
  entirely.
- This repo's own layered structure (`handler.go`/`service.go`/
  `repository.go` per bounded concept) is **flatter** than a Flutter
  Clean Architecture `domain/data/presentation` split — don't reflexively
  rebuild that folder tree here.

## Exercise (Level 4 — Refactor)
`Store.List` has a `// TODO` gap: find the task matching `cursor`, start
from the one right after it (or index 0 if `cursor` is empty/not found),
take up to `limit` tasks, and set `NextCursor` to the last task's ID if
more remain, or `""` at the end. Remove the `t.Skip(...)` calls in
`pagination_test.go` once implemented.
