package restserver

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/sdkopen/sdkopen-go/validator"
)

func init() {
	validator.Initialize()
}

func newTestContext(method, target string, body string) (*chiWebContext, *httptest.ResponseRecorder) {
	var req *http.Request
	if body != "" {
		req = httptest.NewRequest(method, target, strings.NewReader(body))
	} else {
		req = httptest.NewRequest(method, target, nil)
	}
	rec := httptest.NewRecorder()
	return &chiWebContext{writer: rec, request: req}, rec
}

func TestChiWebContext_Context(t *testing.T) {
	ctx, _ := newTestContext("GET", "/", "")
	if ctx.Context() == nil {
		t.Fatal("expected non-nil context")
	}
}

func TestChiWebContext_Response(t *testing.T) {
	ctx, rec := newTestContext("GET", "/", "")
	if ctx.Response() != rec {
		t.Fatal("expected Response() to return the recorder")
	}
}

func TestChiWebContext_Request(t *testing.T) {
	ctx, _ := newTestContext("GET", "/", "")
	if ctx.Request() == nil {
		t.Fatal("expected non-nil request")
	}
}

func TestChiWebContext_RequestHeader(t *testing.T) {
	ctx, _ := newTestContext("GET", "/", "")
	ctx.request.Header.Set("X-Custom", "test-value")

	values := ctx.RequestHeader("X-Custom")
	if len(values) != 1 || values[0] != "test-value" {
		t.Fatalf("expected [test-value], got %v", values)
	}
}

func TestChiWebContext_RequestHeaders(t *testing.T) {
	ctx, _ := newTestContext("GET", "/", "")
	ctx.request.Header.Set("X-First", "one")
	ctx.request.Header.Set("X-Second", "two")

	headers := ctx.RequestHeaders()
	if headers["X-First"][0] != "one" {
		t.Fatalf("expected X-First=one, got %v", headers["X-First"])
	}
	if headers["X-Second"][0] != "two" {
		t.Fatalf("expected X-Second=two, got %v", headers["X-Second"])
	}
}

func TestChiWebContext_RawQuery(t *testing.T) {
	ctx, _ := newTestContext("GET", "/test?foo=bar&baz=qux", "")
	result := ctx.RawQuery()
	if result != "foo=bar&baz=qux" {
		t.Fatalf("expected 'foo=bar&baz=qux', got '%s'", result)
	}
}

func TestChiWebContext_QueryParam(t *testing.T) {
	ctx, _ := newTestContext("GET", "/test?name=John&age=30", "")

	name := ctx.QueryParam("name")
	if name != "John" {
		t.Fatalf("expected name=John, got %s", name)
	}

	age := ctx.QueryParam("age")
	if age != "30" {
		t.Fatalf("expected age=30, got %s", age)
	}
}

func TestChiWebContext_QueryParam_Missing(t *testing.T) {
	ctx, _ := newTestContext("GET", "/test", "")
	result := ctx.QueryParam("missing")
	if result != "" {
		t.Fatalf("expected empty string, got '%s'", result)
	}
}

func TestChiWebContext_QueryArrayParam(t *testing.T) {
	ctx, _ := newTestContext("GET", "/test?tags=a,b,c", "")
	result := ctx.QueryArrayParam("tags")
	if len(result) != 3 {
		t.Fatalf("expected 3 items, got %d: %v", len(result), result)
	}
	if result[0] != "a" || result[1] != "b" || result[2] != "c" {
		t.Fatalf("expected [a b c], got %v", result)
	}
}

func TestChiWebContext_QueryArrayParam_MultipleKeys(t *testing.T) {
	ctx, _ := newTestContext("GET", "/test?tags=a&tags=b,c", "")
	result := ctx.QueryArrayParam("tags")
	if len(result) != 3 {
		t.Fatalf("expected 3 items, got %d: %v", len(result), result)
	}
}

func TestChiWebContext_Path(t *testing.T) {
	ctx, _ := newTestContext("GET", "/api/users", "")
	if ctx.Path() != "/api/users" {
		t.Fatalf("expected /api/users, got %s", ctx.Path())
	}
}

func TestChiWebContext_Method(t *testing.T) {
	ctx, _ := newTestContext("POST", "/", "")
	if ctx.Method() != "POST" {
		t.Fatalf("expected POST, got %s", ctx.Method())
	}
}

func TestChiWebContext_StringBody(t *testing.T) {
	ctx, _ := newTestContext("POST", "/", "hello body")
	body, err := ctx.StringBody()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if body != "hello body" {
		t.Fatalf("expected 'hello body', got '%s'", body)
	}
}

func TestChiWebContext_DecodeBody(t *testing.T) {
	type payload struct {
		Name string `json:"name" validate:"required"`
	}

	data, _ := json.Marshal(payload{Name: "test"})
	ctx, _ := newTestContext("POST", "/", string(data))
	ctx.request.Header.Set("Content-Type", "application/json")

	var result payload
	err := ctx.DecodeBody(&result)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result.Name != "test" {
		t.Fatalf("expected Name=test, got %s", result.Name)
	}
}

func TestChiWebContext_DecodeBody_InvalidJSON(t *testing.T) {
	ctx, _ := newTestContext("POST", "/", "not json")
	ctx.request.Header.Set("Content-Type", "application/json")

	var result map[string]any
	err := ctx.DecodeBody(&result)
	if err == nil {
		t.Fatal("expected error for invalid JSON body, got nil")
	}
}

func TestChiWebContext_DecodeBody_ValidationFails(t *testing.T) {
	type payload struct {
		Name string `json:"name" validate:"required"`
	}

	data, _ := json.Marshal(map[string]string{"name": ""})
	ctx, _ := newTestContext("POST", "/", string(data))

	var result payload
	err := ctx.DecodeBody(&result)
	if err == nil {
		t.Fatal("expected validation error for empty required field, got nil")
	}
}

func TestChiWebContext_AddHeader(t *testing.T) {
	ctx, rec := newTestContext("GET", "/", "")
	ctx.AddHeader("X-Custom", "myvalue")

	if rec.Header().Get("X-Custom") != "myvalue" {
		t.Fatalf("expected X-Custom=myvalue, got %s", rec.Header().Get("X-Custom"))
	}
}

func TestChiWebContext_AddHeaders(t *testing.T) {
	ctx, rec := newTestContext("GET", "/", "")
	ctx.AddHeaders(map[string]string{
		"X-First":  "one",
		"X-Second": "two",
	})

	if rec.Header().Get("X-First") != "one" {
		t.Fatalf("expected X-First=one, got %s", rec.Header().Get("X-First"))
	}
	if rec.Header().Get("X-Second") != "two" {
		t.Fatalf("expected X-Second=two, got %s", rec.Header().Get("X-Second"))
	}
}

func TestChiWebContext_JsonResponse(t *testing.T) {
	ctx, rec := newTestContext("GET", "/", "")
	ctx.JsonResponse(StatusOK, map[string]string{"status": "ok"})

	if rec.Code != 200 {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}
	if ct := rec.Header().Get("Content-Type"); ct != "application/json" {
		t.Fatalf("expected Content-Type application/json, got %s", ct)
	}

	var body map[string]string
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("expected valid JSON response body, got error: %v", err)
	}
	if body["status"] != "ok" {
		t.Fatalf("expected status=ok, got %s", body["status"])
	}
}

func TestChiWebContext_EmptyResponse(t *testing.T) {
	ctx, rec := newTestContext("DELETE", "/", "")
	ctx.EmptyResponse(StatusNoContent)

	if rec.Code != 204 {
		t.Fatalf("expected status 204, got %d", rec.Code)
	}
	if rec.Body.Len() != 0 {
		t.Fatalf("expected empty body, got %s", rec.Body.String())
	}
}

func TestChiWebContext_ErrorResponse(t *testing.T) {
	ctx, rec := newTestContext("GET", "/test", "")
	ctx.ErrorResponse(StatusBadRequest, errors.New("bad input"))

	if rec.Code != 400 {
		t.Fatalf("expected status 400, got %d", rec.Code)
	}

	var body string
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("expected valid JSON, got error: %v", err)
	}
	if body != "bad input" {
		t.Fatalf("expected 'bad input', got '%s'", body)
	}
}
