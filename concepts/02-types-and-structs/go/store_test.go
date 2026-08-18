package tasks

import "testing"

// TestGetOnZeroValueStore proves reading from a nil map is safe — no
// constructor needed, no panic, just a false "found."
func TestGetOnZeroValueStore(t *testing.T) {
	var s Store // zero value: s.tasks is nil

	_, ok := s.Get("missing")
	if ok {
		t.Errorf("Get() on empty store = ok:true, want ok:false")
	}
}

// TestTitlesOnZeroValueStore proves append on a nil slice is safe —
// contrast this with TestAddOnZeroValueStore below.
func TestTitlesOnZeroValueStore(t *testing.T) {
	var s Store

	got := s.Titles()
	if len(got) != 0 {
		t.Errorf("Titles() on empty store = %v, want empty", got)
	}
}

// TestAddOnZeroValueStore is the exercise: Add on a zero-value Store must
// not panic. Skipped until the nil-map guard TODO in store.go is filled in.
func TestAddOnZeroValueStore(t *testing.T) {
	t.Skip("TODO(exercise): guard the nil map in Add, then remove this Skip")

	var s Store // zero value: s.tasks is nil — Add must handle this
	s.Add(Task{ID: "1", Title: "learn Go maps"})

	got, ok := s.Get("1")
	if !ok || got.Title != "learn Go maps" {
		t.Errorf("Get(\"1\") = %+v, ok:%v, want {1 learn Go maps false}, ok:true", got, ok)
	}
}
