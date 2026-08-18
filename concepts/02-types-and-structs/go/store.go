// Package tasks holds Stage 2 exercises: zero values, and the nil-map vs
// nil-slice footgun that has no Dart analogue (Dart's Map/List are never
// "unset," just empty).
package tasks

// Task is a plain struct — its zero value (Task{}) is already fully usable,
// unlike a Dart class with late/nullable fields that must be assigned before
// use.
type Task struct {
	ID    string
	Title string
	Done  bool
}

// Store holds tasks in memory, keyed by ID.
//
// TODO(exercise, Level 2): Store's zero value has a nil map (tasks is
// unexported and never initialized here). Reading from a nil map is safe
// in Go — it just returns the zero value — but WRITING to a nil map
// panics. Fix Add so a zero-value Store (var s Store, or &Store{}) works
// without requiring a separate constructor call first.
type Store struct {
	tasks map[string]Task
}

// Add stores t, keyed by t.ID.
func (s *Store) Add(t Task) {
	// TODO: guard against s.tasks being nil (lazily initialize it here
	// with make(map[string]Task) before the first write), then:
	// s.tasks[t.ID] = t
}

// Get returns the task with the given id, and whether it was found.
// Already correct: reading a nil map is safe and returns (Task{}, false).
func (s *Store) Get(id string) (Task, bool) {
	t, ok := s.tasks[id]
	return t, ok
}

// Titles returns the title of every task, in no particular order.
// Already correct — included to contrast with Store.tasks: appending to a
// nil slice is always safe in Go, unlike writing to a nil map.
func (s *Store) Titles() []string {
	var titles []string // nil slice, not []string{} — and that's fine
	for _, t := range s.tasks {
		titles = append(titles, t.Title) // append on nil works
	}
	return titles
}
