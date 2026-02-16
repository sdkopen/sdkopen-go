package messaging

import "context"

type Publisher interface {
	Publish(ctx context.Context, topic string, body []byte, opts ...PublishOption) error
	Close() error
}

var publisherInstance Publisher

func Publish(ctx context.Context, topic string, body []byte, opts ...PublishOption) error {
	return publisherInstance.Publish(ctx, topic, body, opts...)
}
