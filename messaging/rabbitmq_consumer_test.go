package messaging

import (
	"testing"

	amqp "github.com/rabbitmq/amqp091-go"
)

func TestExtractHeaders_StringValues(t *testing.T) {
	table := amqp.Table{
		"X-Source": "api",
		"X-Trace":  "trace-123",
	}

	headers := extractHeaders(table)

	if len(headers) != 2 {
		t.Fatalf("expected 2 headers, got %d", len(headers))
	}
	if headers["X-Source"] != "api" {
		t.Fatalf("expected X-Source=api, got %s", headers["X-Source"])
	}
	if headers["X-Trace"] != "trace-123" {
		t.Fatalf("expected X-Trace=trace-123, got %s", headers["X-Trace"])
	}
}

func TestExtractHeaders_NonStringValues_Ignored(t *testing.T) {
	table := amqp.Table{
		"string-key": "value",
		"int-key":    int32(42),
		"bool-key":   true,
		"float-key":  3.14,
	}

	headers := extractHeaders(table)

	if len(headers) != 1 {
		t.Fatalf("expected 1 header (only string), got %d", len(headers))
	}
	if headers["string-key"] != "value" {
		t.Fatalf("expected string-key=value, got %s", headers["string-key"])
	}
}

func TestExtractHeaders_EmptyTable(t *testing.T) {
	table := amqp.Table{}
	headers := extractHeaders(table)

	if len(headers) != 0 {
		t.Fatalf("expected 0 headers, got %d", len(headers))
	}
}

func TestExtractHeaders_NilTable(t *testing.T) {
	headers := extractHeaders(nil)

	if headers == nil {
		t.Fatal("expected non-nil map")
	}
	if len(headers) != 0 {
		t.Fatalf("expected 0 headers, got %d", len(headers))
	}
}
