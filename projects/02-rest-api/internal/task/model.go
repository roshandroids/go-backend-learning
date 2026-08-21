// Package task holds Project 2's REST API for a single resource: Task.
// Flat handler/service/repository split, no domain/usecase/entity
// layers — per the roadmap's own "Do NOT copy" note for this project.
package task

// Task is the resource this API serves.
type Task struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// CreateRequest is the JSON body for POST /tasks.
type CreateRequest struct {
	Title string `json:"title"`
}

// Page is one page of List results plus the cursor for the next one.
// Cursor-based, same idiom as concepts/08-rest-api-design, applied here.
type Page struct {
	Tasks      []Task `json:"tasks"`
	NextCursor string `json:"next_cursor,omitempty"`
}

// APIError is the one structured error shape every handler returns —
// decided once, never deviated from, per Stage 6's own guidance.
type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e APIError) Error() string { return e.Message }
