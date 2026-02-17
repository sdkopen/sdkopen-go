package messaging

import (
	"github.com/sdkopen/sdkopen-go/logging"
	"github.com/sdkopen/sdkopen-go/observer"
)

type Provider struct {
	CreatePublisher func() Publisher
	CreateConsumer  func() Consumer
}

func Initialize(provider *Provider) {
	publisherInstance = provider.CreatePublisher()
	observer.Attach(publisherObserver{})
	logging.Info("messaging publisher initialized")

	consumerInstance = provider.CreateConsumer()
	observer.Attach(consumerObserver{})
	logging.Info("messaging consumer initialized")
}
