package messaging

import (
	"context"

	"github.com/sdkopen/sdkopen-go/logging"
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
	subscriptions    []Subscription
)

func Subscribe(topic string, handler HandlerFunc) {
	subscriptions = append(subscriptions, Subscription{topic, handler})
}

func StartConsumer() {
	for _, sub := range subscriptions {
		consumerInstance.Subscribe(sub)
	}
	logging.Info("messaging consumer started")
	consumerInstance.Start()
}
