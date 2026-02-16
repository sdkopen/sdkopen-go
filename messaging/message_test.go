package messaging

import (
	"testing"
)

func TestWithHeaders(t *testing.T) {
	headers := map[string]string{
		"X-Custom": "value",
		"X-Trace":  "abc123",
	}

	opt := WithHeaders(headers)
	cfg := publishConfig{}
	opt(&cfg)

	if len(cfg.Headers) != 2 {
		t.Fatalf("expected 2 headers, got %d", len(cfg.Headers))
	}
	if cfg.Headers["X-Custom"] != "value" {
		t.Fatalf("expected X-Custom=value, got %s", cfg.Headers["X-Custom"])
	}
	if cfg.Headers["X-Trace"] != "abc123" {
		t.Fatalf("expected X-Trace=abc123, got %s", cfg.Headers["X-Trace"])
	}
}

func TestWithDelay(t *testing.T) {
	opt := WithDelay(30)
	cfg := publishConfig{}
	opt(&cfg)

	if cfg.DelaySeconds != 30 {
		t.Fatalf("expected DelaySeconds=30, got %d", cfg.DelaySeconds)
	}
}

func TestWithDelay_Zero(t *testing.T) {
	opt := WithDelay(0)
	cfg := publishConfig{}
	opt(&cfg)

	if cfg.DelaySeconds != 0 {
		t.Fatalf("expected DelaySeconds=0, got %d", cfg.DelaySeconds)
	}
}

func TestApplyOptions_NoOptions(t *testing.T) {
	cfg := applyOptions(nil)

	if cfg.Headers != nil {
		t.Fatalf("expected nil headers, got %v", cfg.Headers)
	}
	if cfg.DelaySeconds != 0 {
		t.Fatalf("expected DelaySeconds=0, got %d", cfg.DelaySeconds)
	}
}

func TestApplyOptions_MultipleOptions(t *testing.T) {
	headers := map[string]string{"key": "val"}
	opts := []PublishOption{
		WithHeaders(headers),
		WithDelay(10),
	}

	cfg := applyOptions(opts)

	if cfg.Headers["key"] != "val" {
		t.Fatalf("expected key=val, got %s", cfg.Headers["key"])
	}
	if cfg.DelaySeconds != 10 {
		t.Fatalf("expected DelaySeconds=10, got %d", cfg.DelaySeconds)
	}
}

func TestApplyOptions_LastOptionWins(t *testing.T) {
	opts := []PublishOption{
		WithDelay(5),
		WithDelay(20),
	}

	cfg := applyOptions(opts)

	if cfg.DelaySeconds != 20 {
		t.Fatalf("expected DelaySeconds=20 (last option wins), got %d", cfg.DelaySeconds)
	}
}

func TestMessage_Struct(t *testing.T) {
	msg := Message{
		ID:    "msg-123",
		Topic: "orders",
		Body:  []byte(`{"order_id":1}`),
		Headers: map[string]string{
			"source": "test",
		},
	}

	if msg.ID != "msg-123" {
		t.Fatalf("expected ID=msg-123, got %s", msg.ID)
	}
	if msg.Topic != "orders" {
		t.Fatalf("expected Topic=orders, got %s", msg.Topic)
	}
	if string(msg.Body) != `{"order_id":1}` {
		t.Fatalf("expected body={\"order_id\":1}, got %s", string(msg.Body))
	}
	if msg.Headers["source"] != "test" {
		t.Fatalf("expected source=test, got %s", msg.Headers["source"])
	}
}
