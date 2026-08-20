// Package httpserver holds Stage 5 exercises: net/http handlers and
// middleware as plain function composition — no framework, no
// decorators, no annotations.
package httpserver

import (
	"encoding/json"
	"net/http"
)

// LoggingMiddleware wraps a handler with structured request logging.
// This IS the whole idiom: middleware is a function that takes a
// http.Handler and returns a http.Handler. Every third-party
// router/middleware library you'll meet later is built from this same
// shape — there's no hidden framework mechanism to learn.
//
// TODO(exercise, Level 2): record time.Now(), call next.ServeHTTP(w, r),
// then log the method, path, and elapsed duration via
// slog.Info("request", "method", r.Method, "path", r.URL.Path,
// "duration", time.Since(start)).
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: start := time.Now()
		next.ServeHTTP(w, r)
		// TODO: slog.Info("request", "method", r.Method, "path", r.URL.Path, "duration", time.Since(start))
	})
}

// HealthHandler responds 200 with a small JSON body. Already implemented
// — JSON encoding is a "skim" topic (conceptually identical to
// toJson/fromJson); the point of this module is the middleware above,
// not this handler.
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

// User is the JSON shape for the two handlers below.
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// GetUserHandler reads {id} from the request path (Go 1.22+'s
// ServeMux pattern matching, no router library needed) and echoes it
// back as JSON.
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(User{ID: id, Name: "placeholder"})
}

// CreateUserHandler decodes a JSON request body into a User and echoes
// it back — the mirror image of every Dio/http call you've made as a
// Flutter client, now from the server side.
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}

// NewMux wires the three routes together. No router library — stdlib
// http.ServeMux only, per the Stage 5 exercise's own constraint.
func NewMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", HealthHandler)
	mux.HandleFunc("GET /users/{id}", GetUserHandler)
	mux.HandleFunc("POST /users", CreateUserHandler)
	return mux
}
