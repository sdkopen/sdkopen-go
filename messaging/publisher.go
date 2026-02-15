package messaging

import (
	"context"

	"github.com/sdkopen/sdkopen-go/logging"
	"github.com/sdkopen/sdkopen-go/observer"
)

type Publisher interface {
	Publish(ctx context.Context, topic string, body []byte, opts ...PublishOption) error
	Close() error
}

var publisherInstance Publisher

func InitializePublisher(factory func() Publisher) {
	publisherInstance = factory()
	observer.Attach(publisherObserver{})
	logging.Info("messaging publisher initialized")
}

func Publish(ctx context.Context, topic string, body []byte, opts ...PublishOption) error {
	return publisherInstance.Publish(ctx, topic, body, opts...)
}
