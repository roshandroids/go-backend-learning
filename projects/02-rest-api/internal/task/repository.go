package task

import "sync"

// Repository is an in-memory, insertion-ordered Task store, guarded by
// a mutex because — unlike the single-threaded concept-module version —
// a real HTTP server serves concurrent requests. This is Stage 4's
// lesson (shared mutable state needs synchronization) applied here, not
// re-taught.
type Repository struct {
	mu     sync.RWMutex
	order  []string // insertion-ordered IDs; doubles as List's cursor space
	tasks  map[string]Task
	nextID int
}

func NewRepository() *Repository {
	return &Repository{tasks: make(map[string]Task)}
}

// Create assigns the next ID, stores the task, and returns it.
//
// TODO(exercise, Level 6 — Build): increment r.nextID, format it as the
// task's ID (strconv.Itoa), build a Task{ID, Title: title}, store it in
// r.tasks, append its ID to r.order, and return the Task.
func (r *Repository) Create(title string) Task {
	r.mu.Lock()
	defer r.mu.Unlock()
	// TODO
	return Task{}
}

// Get returns the task with the given id, and whether it was found.
//
// TODO(exercise, Level 2 — Complete): look up id in r.tasks and return
// it along with the found bool.
func (r *Repository) Get(id string) (Task, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	// TODO
	return Task{}, false
}

// List returns up to limit tasks starting immediately after the task
// whose ID equals cursor — the exact same cursor-pagination algorithm
// as concepts/08-rest-api-design's Store.List, applied to a real
// resource here rather than introduced fresh.
//
// TODO(exercise, Level 4 — Refactor): find the index of cursor in
// r.order (or start at 0 if cursor is empty/not found), take up to
// limit IDs from there, look each one up in r.tasks, and set
// NextCursor to the last included ID if more remain, or "" at the end.
func (r *Repository) List(cursor string, limit int) Page {
	r.mu.RLock()
	defer r.mu.RUnlock()
	// TODO
	return Page{}
}
