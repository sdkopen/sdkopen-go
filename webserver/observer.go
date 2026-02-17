package webserver

import (
	"github.com/sdkopen/sdkopen-go/common/observer"
	"github.com/sdkopen/sdkopen-go/logging"
)

type webServerObserver struct {
}

func (o webServerObserver) Close() {
	logging.Info("waiting to safely close the http server")
	if observer.WaitRunningTimeout() {
		logging.Warn("WaitGroup timed out, forcing close http server")
	}
	logging.Info("closing http server")
	if err := ServerInstance.Shutdown(); err != nil {
		logging.Error("error when closing http server: %v", err)
	}
	ServerInstance = nil
}
