package messaging

import (
	"context"
)

type HandlerFunc func(ctx context.Context, msg Message) error

type Subscription struct {
	Topic   string
	Handler HandlerFunc
}

type Consumer interface {
	Subscribe(subscription Subscription)
	Start() error
	Close() error
}

var (
	consumerInstance Consumer
)

func Subscribe(topic string, handler HandlerFunc) {
	consumerInstance.Subscribe(Subscription{topic, handler})
}
