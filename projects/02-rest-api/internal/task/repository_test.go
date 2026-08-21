package task

import "testing"

// TestRepositoryCreateAndGet is the exercise's happy path. Skipped
// until the TODO in repository.go is filled in.
func TestRepositoryCreateAndGet(t *testing.T) {
	t.Skip("TODO(exercise): implement Repository.Create and Get, then remove this Skip")

	r := NewRepository()
	created := r.Create("buy milk")

	if created.ID == "" {
		t.Fatal("Create() returned a Task with an empty ID")
	}
	if created.Title != "buy milk" {
		t.Errorf("Create().Title = %q, want %q", created.Title, "buy milk")
	}

	got, ok := r.Get(created.ID)
	if !ok {
		t.Fatalf("Get(%q) ok = false, want true", created.ID)
	}
	if got != created {
		t.Errorf("Get(%q) = %+v, want %+v", created.ID, got, created)
	}
}

// TestRepositoryGetMissingReturnsFalse doesn't depend on Create.
// Skipped anyway so the exercise isn't "partially done" by accident —
// Get itself still has a TODO.
func TestRepositoryGetMissingReturnsFalse(t *testing.T) {
	t.Skip("TODO(exercise): implement Repository.Get, then remove this Skip")

	r := NewRepository()
	_, ok := r.Get("does-not-exist")
	if ok {
		t.Error("Get() on empty repository = ok:true, want ok:false")
	}
}

// TestRepositoryListPaginatesInInsertionOrder is the same cursor-
// pagination exercise as concepts/08-rest-api-design, applied here.
// Skipped until the TODOs in repository.go are filled in.
func TestRepositoryListPaginatesInInsertionOrder(t *testing.T) {
	t.Skip("TODO(exercise): implement Repository.Create and List, then remove this Skip")

	r := NewRepository()
	var created []Task
	for i := 0; i < 5; i++ {
		created = append(created, r.Create("task"))
	}

	page := r.List("", 2)
	if len(page.Tasks) != 2 {
		t.Fatalf("got %d tasks, want 2", len(page.Tasks))
	}
	if page.Tasks[0].ID != created[0].ID || page.Tasks[1].ID != created[1].ID {
		t.Errorf("page.Tasks = %+v, want first two created tasks in order", page.Tasks)
	}
	if page.NextCursor != created[1].ID {
		t.Errorf("NextCursor = %q, want %q", page.NextCursor, created[1].ID)
	}

	last := r.List(created[3].ID, 2)
	if len(last.Tasks) != 1 || last.Tasks[0].ID != created[4].ID {
		t.Errorf("last page = %+v, want just the final created task", last.Tasks)
	}
	if last.NextCursor != "" {
		t.Errorf("NextCursor on last page = %q, want empty", last.NextCursor)
	}
}
