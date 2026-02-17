package commonhttp

import "testing"

func TestHttpStatusCode_Int(t *testing.T) {
	tests := []struct {
		code     HttpStatusCode
		expected int
	}{
		{StatusOK, 200},
		{StatusCreated, 201},
		{StatusNoContent, 204},
		{StatusBadRequest, 400},
		{StatusUnauthorized, 401},
		{StatusForbidden, 403},
		{StatusNotFound, 404},
		{StatusInternalServerError, 500},
	}

	for _, tt := range tests {
		t.Run(tt.code.String(), func(t *testing.T) {
			result := tt.code.Int()
			if result != tt.expected {
				t.Fatalf("expected %d, got %d", tt.expected, result)
			}
		})
	}
}

func TestHttpStatusCode_String(t *testing.T) {
	tests := []struct {
		code     HttpStatusCode
		expected string
	}{
		{StatusContinue, "Continue"},
		{StatusOK, "OK"},
		{StatusCreated, "Created"},
		{StatusNoContent, "No Content"},
		{StatusMovedPermanently, "Moved Permanently"},
		{StatusBadRequest, "Bad Request"},
		{StatusUnauthorized, "Unauthorized"},
		{StatusForbidden, "Forbidden"},
		{StatusNotFound, "Not Found"},
		{StatusMethodNotAllowed, "Method Not Allowed"},
		{StatusConflict, "Conflict"},
		{StatusUnprocessableEntity, "Unprocessable Entity"},
		{StatusTooManyRequests, "Too Many Requests"},
		{StatusInternalServerError, "Internal Server Error"},
		{StatusServiceUnavailable, "Service Unavailable"},
		{StatusTeapot, "Teapot"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := tt.code.String()
			if result != tt.expected {
				t.Fatalf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestHttpStatusCode_String_UnknownCode(t *testing.T) {
	code := HttpStatusCode(999)
	result := code.String()
	if result != "" {
		t.Fatalf("expected empty string for unknown code, got %s", result)
	}
}
