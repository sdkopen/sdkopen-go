package commonhttp

import "testing"

func TestContentType_String(t *testing.T) {
	tests := []struct {
		ct       ContentType
		expected string
	}{
		{ContentTypeJSON, "application/json"},
		{ContentTypeTextPlain, "text/plain"},
		{ContentTypePDF, "application/pdf"},
		{ContentTypeOctetStream, "application/octet-stream"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := tt.ct.String()
			if result != tt.expected {
				t.Fatalf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}
