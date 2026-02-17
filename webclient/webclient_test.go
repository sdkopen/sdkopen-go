package webclient

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	commonhttp "github.com/sdkopen/sdkopen-go/common/http"
)

func TestNew(t *testing.T) {
	client := New("https://api.example.com")
	if client.baseURL != "https://api.example.com" {
		t.Fatalf("expected baseURL https://api.example.com, got %s", client.baseURL)
	}
}

func TestWithHeader(t *testing.T) {
	client := New("https://api.example.com").
		WithHeader("Authorization", "Bearer token")

	if client.headers["Authorization"] != "Bearer token" {
		t.Fatalf("expected Authorization header, got %v", client.headers)
	}
}

func TestWithTimeout(t *testing.T) {
	client := New("https://api.example.com").
		WithTimeout(5 * time.Second)

	if client.client.Timeout != 5*time.Second {
		t.Fatalf("expected 5s timeout, got %v", client.client.Timeout)
	}
}

func TestGet(t *testing.T) {
	type user struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/users/1" {
			t.Fatalf("expected /users/1, got %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user{ID: 1, Name: "Alice"})
	}))
	defer server.Close()

	client := New(server.URL)
	var result user
	resp, err := client.Get("/users/1", &result)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if resp.StatusCode != commonhttp.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if result.Name != "Alice" {
		t.Fatalf("expected Alice, got %s", result.Name)
	}
}

func TestPost(t *testing.T) {
	type user struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("expected POST, got %s", r.Method)
		}
		if ct := r.Header.Get("Content-Type"); ct != "application/json" {
			t.Fatalf("expected Content-Type application/json, got %s", ct)
		}

		var body user
		json.NewDecoder(r.Body).Decode(&body)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user{ID: 1, Name: body.Name})
	}))
	defer server.Close()

	client := New(server.URL)
	var result user
	resp, err := client.Post("/users", user{Name: "Bob"}, &result)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if resp.StatusCode != commonhttp.StatusCreated {
		t.Fatalf("expected 201, got %d", resp.StatusCode)
	}
	if result.Name != "Bob" {
		t.Fatalf("expected Bob, got %s", result.Name)
	}
}

func TestPut(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Fatalf("expected PUT, got %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"updated"}`))
	}))
	defer server.Close()

	client := New(server.URL)
	var result map[string]string
	resp, err := client.Put("/users/1", map[string]string{"name": "Charlie"}, &result)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if resp.StatusCode != commonhttp.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if result["status"] != "updated" {
		t.Fatalf("expected updated, got %s", result["status"])
	}
}

func TestPatch(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Fatalf("expected PATCH, got %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"patched"}`))
	}))
	defer server.Close()

	client := New(server.URL)
	var result map[string]string
	resp, err := client.Patch("/users/1", map[string]string{"name": "Dave"}, &result)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if resp.StatusCode != commonhttp.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if result["status"] != "patched" {
		t.Fatalf("expected patched, got %s", result["status"])
	}
}

func TestDelete(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Fatalf("expected DELETE, got %s", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	client := New(server.URL)
	resp, err := client.Delete("/users/1", nil)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if resp.StatusCode != commonhttp.StatusNoContent {
		t.Fatalf("expected 204, got %d", resp.StatusCode)
	}
}

func TestWithHeaders_SentInRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if auth := r.Header.Get("Authorization"); auth != "Bearer mytoken" {
			t.Fatalf("expected Authorization Bearer mytoken, got %s", auth)
		}
		if custom := r.Header.Get("X-Custom"); custom != "value" {
			t.Fatalf("expected X-Custom value, got %s", custom)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := New(server.URL).
		WithHeader("Authorization", "Bearer mytoken").
		WithHeader("X-Custom", "value")

	resp, err := client.Get("/test", nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if resp.StatusCode != commonhttp.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
}

func TestGet_NilResult(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"data":"value"}`))
	}))
	defer server.Close()

	client := New(server.URL)
	resp, err := client.Get("/test", nil)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if resp.StatusCode != commonhttp.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
}

func TestResponse_Body(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`raw body content`))
	}))
	defer server.Close()

	client := New(server.URL)
	resp, err := client.Get("/test", nil)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if string(resp.Body) != "raw body content" {
		t.Fatalf("expected 'raw body content', got '%s'", string(resp.Body))
	}
}

func TestResponse_Headers(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Response-Id", "abc123")
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := New(server.URL)
	resp, err := client.Get("/test", nil)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if resp.Headers.Get("X-Response-Id") != "abc123" {
		t.Fatalf("expected X-Response-Id abc123, got %s", resp.Headers.Get("X-Response-Id"))
	}
}
