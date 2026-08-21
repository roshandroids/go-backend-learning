package task

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func newTestHandler() *Handler {
	return NewHandler(NewService(NewRepository()))
}

// TestCreateHandlerRejectsInvalidJSON doesn't depend on the exercise
// TODOs — json.Decode fails before the service is ever called. Already
// passing.
func TestCreateHandlerRejectsInvalidJSON(t *testing.T) {
	h := newTestHandler()
	req := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader("{not json"))
	w := httptest.NewRecorder()

	h.Create(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("got status %d, want %d", w.Code, http.StatusBadRequest)
	}
}

// TestGetHandlerNotFoundForUnknownID doesn't depend on the exercise
// TODOs — an ID that was never created is not found regardless of
// whether Repository.Get/Create are implemented yet. Already passing.
func TestGetHandlerNotFoundForUnknownID(t *testing.T) {
	h := newTestHandler()
	mux := NewMux(h)
	req := httptest.NewRequest(http.MethodGet, "/tasks/does-not-exist", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("got status %d, want %d", w.Code, http.StatusNotFound)
	}
}

// TestCreateHandlerHappyPath is the exercise: a valid request must
// create and return a Task. Skipped until Repository.Create and
// validateCreateRequest are implemented.
func TestCreateHandlerHappyPath(t *testing.T) {
	t.Skip("TODO(exercise): implement Repository.Create and validateCreateRequest, then remove this Skip")

	h := newTestHandler()
	body := `{"title":"buy milk"}`
	req := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(body))
	w := httptest.NewRecorder()

	h.Create(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("got status %d, want %d", w.Code, http.StatusCreated)
	}
	var got Task
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatalf("decoding response: %v", err)
	}
	if got.Title != "buy milk" || got.ID == "" {
		t.Errorf("got %+v, want a Task with Title \"buy milk\" and a non-empty ID", got)
	}
}

// TestCreateHandlerRejectsEmptyTitle is the structured-error-response
// exercise. Skipped until validateCreateRequest is implemented.
func TestCreateHandlerRejectsEmptyTitle(t *testing.T) {
	t.Skip("TODO(exercise): implement validateCreateRequest, then remove this Skip")

	h := newTestHandler()
	req := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(`{"title":""}`))
	w := httptest.NewRecorder()

	h.Create(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("got status %d, want %d", w.Code, http.StatusBadRequest)
	}
	var got APIError
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatalf("decoding response: %v", err)
	}
	if got.Code != "validation_failed" {
		t.Errorf("got APIError.Code = %q, want %q", got.Code, "validation_failed")
	}
}

// TestListHandlerReturnsCreatedTasks is the exercise, applying
// concepts/08's pagination lesson through the full HTTP stack. Skipped
// until Repository.Create and List are implemented.
func TestListHandlerReturnsCreatedTasks(t *testing.T) {
	t.Skip("TODO(exercise): implement Repository.Create and List, then remove this Skip")

	h := newTestHandler()
	for _, title := range []string{"a", "b", "c"} {
		req := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(`{"title":"`+title+`"}`))
		h.Create(httptest.NewRecorder(), req)
	}

	req := httptest.NewRequest(http.MethodGet, "/tasks?limit=2", nil)
	w := httptest.NewRecorder()
	h.List(w, req)

	var page Page
	if err := json.Unmarshal(w.Body.Bytes(), &page); err != nil {
		t.Fatalf("decoding response: %v", err)
	}
	if len(page.Tasks) != 2 {
		t.Errorf("got %d tasks, want 2", len(page.Tasks))
	}
	if page.NextCursor == "" {
		t.Error("NextCursor is empty, want a cursor since more tasks remain")
	}
}
