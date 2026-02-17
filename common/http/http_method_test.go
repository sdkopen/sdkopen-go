package commonhttp

import "testing"

func TestHttpMethod_String(t *testing.T) {
	tests := []struct {
		method   HttpMethod
		expected string
	}{
		{Get, "GET"},
		{Head, "HEAD"},
		{Post, "POST"},
		{Put, "PUT"},
		{Patch, "PATCH"},
		{Delete, "DELETE"},
		{Connect, "CONNECT"},
		{Options, "OPTIONS"},
		{Trace, "TRACE"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := tt.method.String()
			if result != tt.expected {
				t.Fatalf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestHttpMethod_IotaValues(t *testing.T) {
	if Get != 0 {
		t.Fatalf("expected Get=0, got %d", Get)
	}
	if Trace != 8 {
		t.Fatalf("expected Trace=8, got %d", Trace)
	}
}
