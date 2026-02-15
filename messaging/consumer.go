package messaging

import (
	"context"

	"github.com/sdkopen/sdkopen-go/logging"
	"github.com/sdkopen/sdkopen-go/observer"
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
	subscriptions = append(subscriptions, Subscription{Topic: topic, Handler: handler})
}

func StartConsumer(factory func() Consumer) {
	consumerInstance = factory()
	for _, sub := range subscriptions {
		consumerInstance.Subscribe(sub)
	}
	observer.Attach(consumerObserver{})
	logging.Info("messaging consumer started")
	consumerInstance.Start()
}
