// Package tasks holds Stage 6's cursor-pagination exercise. Structured
// error responses and validation are deliberately not re-taught here —
// concepts/05-error-handling already covers the custom-error-type idiom
// this module would otherwise repeat.
package tasks

// Task is the resource being paginated.
type Task struct {
	ID    string
	Title string
}

// Store is an in-memory, insertion-ordered Task repository.
type Store struct {
	tasks []Task // ordered by insertion; ID doubles as the pagination cursor
}

// Add appends t to the store.
func (s *Store) Add(t Task) {
	s.tasks = append(s.tasks, t)
}

// Page is one page of results plus the cursor for fetching the next one.
// An empty NextCursor means there are no more results.
type Page struct {
	Tasks      []Task
	NextCursor string
}

// List returns up to limit tasks starting immediately after the task
// whose ID equals cursor (an empty cursor starts from the beginning).
//
// This is cursor-based pagination, not offset pagination
// (LIMIT/OFFSET): results stay correct even if tasks are added or
// removed between page requests, and — the reason it matters once
// Stage 7 introduces a real database — it doesn't degrade on large
// tables the way OFFSET does, since OFFSET still has to scan and discard
// every row before it. Same lesson as infinite-scroll pagination in
// Flutter, just enforced server-side instead of client-side.
//
// TODO(exercise, Level 4 — Refactor): find the index of the task
// matching `cursor` and start from the one right after it (or index 0 if
// cursor is empty or not found). Take up to `limit` tasks from there.
// Set NextCursor to the last returned task's ID if more tasks remain
// after this page, or "" if this page reached the end.
func (s *Store) List(cursor string, limit int) Page {
	// TODO
	return Page{}
}
