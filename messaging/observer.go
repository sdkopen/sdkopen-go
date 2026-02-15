package messaging

import (
	"github.com/sdkopen/sdkopen-go/logging"
	"github.com/sdkopen/sdkopen-go/observer"
)

type publisherObserver struct{}

func (o publisherObserver) Close() {
	logging.Info("waiting to safely close the messaging publisher")
	if observer.WaitRunningTimeout() {
		logging.Warn("WaitGroup timed out, forcing close messaging publisher")
	}
	logging.Info("closing messaging publisher")
	if publisherInstance == nil {
		return
	}
	if err := publisherInstance.Close(); err != nil {
		logging.Error("error when closing messaging publisher: %v", err)
	}
	publisherInstance = nil
}

type consumerObserver struct{}

func (o consumerObserver) Close() {
	logging.Info("waiting to safely close the messaging consumer")
	if observer.WaitRunningTimeout() {
		logging.Warn("WaitGroup timed out, forcing close messaging consumer")
	}
	logging.Info("closing messaging consumer")
	if consumerInstance == nil {
		return
	}
	if err := consumerInstance.Close(); err != nil {
		logging.Error("error when closing messaging consumer: %v", err)
	}
	consumerInstance = nil
}
