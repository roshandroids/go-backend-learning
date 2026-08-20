package httpserver

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	HealthHandler(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("got status %d, want %d", w.Code, http.StatusOK)
	}
}

func TestGetUserHandlerUsesPathValue(t *testing.T) {
	mux := NewMux()
	req := httptest.NewRequest(http.MethodGet, "/users/42", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("got status %d, want %d", w.Code, http.StatusOK)
	}
	if !strings.Contains(w.Body.String(), `"id":"42"`) {
		t.Errorf("body = %s, want it to contain \"id\":\"42\"", w.Body.String())
	}
}

func TestCreateUserHandlerDecodesJSON(t *testing.T) {
	body := `{"id":"1","name":"Ada"}`
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(body))
	w := httptest.NewRecorder()

	CreateUserHandler(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("got status %d, want %d", w.Code, http.StatusCreated)
	}
	if !strings.Contains(w.Body.String(), `"name":"Ada"`) {
		t.Errorf("body = %s, want it to contain \"name\":\"Ada\"", w.Body.String())
	}
}

// TestLoggingMiddlewarePassesThroughToNext proves the middleware still
// calls the wrapped handler and leaves its response untouched — the
// property that matters regardless of what logging looks like. Skipped
// until the TODO in server.go is filled in (the placeholder body already
// calls next.ServeHTTP, so this technically compiles and could pass
// early — it's skipped anyway so the exercise isn't "already done" by
// accident).
func TestLoggingMiddlewarePassesThroughToNext(t *testing.T) {
	t.Skip("TODO(exercise): implement the logging in LoggingMiddleware, then remove this Skip")

	called := false
	inner := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(http.StatusTeapot)
	})

	req := httptest.NewRequest(http.MethodGet, "/anything", nil)
	w := httptest.NewRecorder()

	LoggingMiddleware(inner).ServeHTTP(w, req)

	if !called {
		t.Error("LoggingMiddleware did not call the wrapped handler")
	}
	if w.Code != http.StatusTeapot {
		t.Errorf("got status %d, want %d (middleware must not alter the response)", w.Code, http.StatusTeapot)
	}
}
