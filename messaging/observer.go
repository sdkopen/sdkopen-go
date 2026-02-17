package messaging

import (
	"github.com/sdkopen/sdkopen-go/logging"
	"github.com/sdkopen/sdkopen-go/observer"
)

type messagingObserver struct{}

func (o messagingObserver) Close() {
	logging.Info("waiting to safely close the messaging connection")
	if observer.WaitRunningTimeout() {
		logging.Warn("WaitGroup timed out, forcing close messaging connection")
	}

	logging.Info("closing messaging consumer")
	if consumerInstance != nil {
		if err := consumerInstance.Close(); err != nil {
			logging.Error("error when closing messaging consumer: %v", err)
		}
		consumerInstance = nil
	}

	logging.Info("closing messaging publisher")
	if publisherInstance != nil {
		if err := publisherInstance.Close(); err != nil {
			logging.Error("error when closing messaging publisher: %v", err)
		}
		publisherInstance = nil
	}
}
