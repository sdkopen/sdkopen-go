package messaging

import (
	"context"
	"testing"
)

func TestSubscribe_AddsSubscription(t *testing.T) {
	// Reset global state
	subscriptions = nil

	handler := func(ctx context.Context, msg Message) error {
		return nil
	}

	Subscribe("test-topic", handler)

	if len(subscriptions) != 1 {
		t.Fatalf("expected 1 subscription, got %d", len(subscriptions))
	}
	if subscriptions[0].Topic != "test-topic" {
		t.Fatalf("expected topic=test-topic, got %s", subscriptions[0].Topic)
	}
	if subscriptions[0].Handler == nil {
		t.Fatal("expected non-nil handler")
	}
}

func TestSubscribe_MultipleTopics(t *testing.T) {
	subscriptions = nil

	handler := func(ctx context.Context, msg Message) error {
		return nil
	}

	Subscribe("topic-a", handler)
	Subscribe("topic-b", handler)
	Subscribe("topic-c", handler)

	if len(subscriptions) != 3 {
		t.Fatalf("expected 3 subscriptions, got %d", len(subscriptions))
	}
	if subscriptions[0].Topic != "topic-a" {
		t.Fatalf("expected first topic=topic-a, got %s", subscriptions[0].Topic)
	}
	if subscriptions[1].Topic != "topic-b" {
		t.Fatalf("expected second topic=topic-b, got %s", subscriptions[1].Topic)
	}
	if subscriptions[2].Topic != "topic-c" {
		t.Fatalf("expected third topic=topic-c, got %s", subscriptions[2].Topic)
	}
}

func TestSubscription_Struct(t *testing.T) {
	handler := func(ctx context.Context, msg Message) error {
		return nil
	}

	sub := Subscription{
		Topic:   "orders.created",
		Handler: handler,
	}

	if sub.Topic != "orders.created" {
		t.Fatalf("expected Topic=orders.created, got %s", sub.Topic)
	}
	if sub.Handler == nil {
		t.Fatal("expected non-nil Handler")
	}
}
