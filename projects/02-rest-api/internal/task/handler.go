package task

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// Handler is the HTTP layer: decode request, call the service, encode
// response. Fully implemented — JSON/routing are "skim" topics, per
// this repo's own roadmap.
type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

func writeError(w http.ResponseWriter, status int, code, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(APIError{Code: code, Message: message})
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid_json", "request body must be valid JSON")
		return
	}

	t, err := h.svc.Create(req)
	if err != nil {
		writeError(w, http.StatusBadRequest, "validation_failed", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(t)
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	t, ok := h.svc.Get(id)
	if !ok {
		writeError(w, http.StatusNotFound, "not_found", "task not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(t)
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	cursor := r.URL.Query().Get("cursor")
	limit := 20
	if v := r.URL.Query().Get("limit"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			limit = n
		}
	}

	page := h.svc.List(cursor, limit)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(page)
}

// NewMux wires the three routes. Stdlib http.ServeMux only, no router
// library, matching Stage 5's own constraint.
func NewMux(h *Handler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /tasks", h.Create)
	mux.HandleFunc("GET /tasks/{id}", h.Get)
	mux.HandleFunc("GET /tasks", h.List)
	return mux
}
