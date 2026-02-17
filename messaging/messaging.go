package messaging

import (
	"github.com/sdkopen/sdkopen-go/common/observer"
	"github.com/sdkopen/sdkopen-go/logging"
)

type Provider struct {
	CreatePublisher func() Publisher
	CreateConsumer  func() Consumer
}

func Initialize(provider *Provider) {
	publisherInstance = provider.CreatePublisher()
	logging.Info("messaging publisher initialized")

	consumerInstance = provider.CreateConsumer()
	logging.Info("messaging consumer initialized")

	if err := observer.Attach(messagingObserver{}); err != nil {
		logging.Fatal("could not attach messaging to observer: %v", err)
		return
	}
	logging.Info("messaging connected")
}
