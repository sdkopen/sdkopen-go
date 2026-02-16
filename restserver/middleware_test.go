package restserver

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAccessControlMiddleware_SetsCORSHeaders(t *testing.T) {
	handler := accessControlMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if origin := rec.Header().Get("Access-Control-Allow-Origin"); origin != "*" {
		t.Fatalf("expected Access-Control-Allow-Origin=*, got %s", origin)
	}

	methods := rec.Header().Get("Access-Control-Allow-Methods")
	if methods != "OPTIONS, GET, POST, PUT, PATCH, DELETE" {
		t.Fatalf("expected CORS methods, got %s", methods)
	}

	headers := rec.Header().Get("Access-Control-Allow-Headers")
	if headers != "*" {
		t.Fatalf("expected Access-Control-Allow-Headers=*, got %s", headers)
	}
}

func TestAccessControlMiddleware_OptionsRequest_ShortCircuits(t *testing.T) {
	nextCalled := false
	handler := accessControlMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextCalled = true
	}))

	req := httptest.NewRequest(http.MethodOptions, "/test", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if nextCalled {
		t.Fatal("expected OPTIONS request to short-circuit, but next handler was called")
	}
}

func TestAccessControlMiddleware_NonOptionsRequest_CallsNext(t *testing.T) {
	nextCalled := false
	handler := accessControlMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextCalled = true
	}))

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if !nextCalled {
		t.Fatal("expected GET request to call next handler")
	}
}
