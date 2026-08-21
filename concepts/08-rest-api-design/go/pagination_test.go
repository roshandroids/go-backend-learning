package tasks

import (
	"reflect"
	"testing"
)

func newFiveTaskStore() *Store {
	s := &Store{}
	for i := 1; i <= 5; i++ {
		id := string(rune('0' + i))
		s.Add(Task{ID: id, Title: "task " + id})
	}
	return s
}

// TestListFirstPage is the exercise: an empty cursor starts from the
// beginning. Skipped until the TODO in pagination.go is filled in.
func TestListFirstPage(t *testing.T) {
	t.Skip("TODO(exercise): implement Store.List, then remove this Skip")

	s := newFiveTaskStore()
	page := s.List("", 2)

	wantIDs := []string{"1", "2"}
	if got := ids(page.Tasks); !reflect.DeepEqual(got, wantIDs) {
		t.Errorf("List(\"\", 2).Tasks IDs = %v, want %v", got, wantIDs)
	}
	if page.NextCursor != "2" {
		t.Errorf("NextCursor = %q, want %q", page.NextCursor, "2")
	}
}

// TestListMiddlePage proves paging continues from a non-empty cursor.
// Skipped until the TODO in pagination.go is filled in.
func TestListMiddlePage(t *testing.T) {
	t.Skip("TODO(exercise): implement Store.List, then remove this Skip")

	s := newFiveTaskStore()
	page := s.List("2", 2)

	wantIDs := []string{"3", "4"}
	if got := ids(page.Tasks); !reflect.DeepEqual(got, wantIDs) {
		t.Errorf("List(\"2\", 2).Tasks IDs = %v, want %v", got, wantIDs)
	}
	if page.NextCursor != "4" {
		t.Errorf("NextCursor = %q, want %q", page.NextCursor, "4")
	}
}

// TestListLastPageHasNoNextCursor proves the terminal page signals "no
// more" with an empty NextCursor. Skipped until the TODO is filled in.
func TestListLastPageHasNoNextCursor(t *testing.T) {
	t.Skip("TODO(exercise): implement Store.List, then remove this Skip")

	s := newFiveTaskStore()
	page := s.List("4", 2)

	wantIDs := []string{"5"}
	if got := ids(page.Tasks); !reflect.DeepEqual(got, wantIDs) {
		t.Errorf("List(\"4\", 2).Tasks IDs = %v, want %v", got, wantIDs)
	}
	if page.NextCursor != "" {
		t.Errorf("NextCursor = %q, want empty (no more pages)", page.NextCursor)
	}
}

func ids(tasks []Task) []string {
	out := make([]string, len(tasks))
	for i, t := range tasks {
		out[i] = t.ID
	}
	return out
}
