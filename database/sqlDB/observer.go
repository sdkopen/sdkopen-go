package sqlDB

import (
	"database/sql"

	"github.com/sdkopen/sdkopen-go/logging"
	"github.com/sdkopen/sdkopen-go/observer"
)

const (
	observerWaitingMsg      string = "Waiting to safely close the %s database connection"
	observerForcingMsg      string = "WaitGroup timed out, forcing close the %s database connection"
	observerClosingMsg      string = "Closing %s database connection"
	observerClosingErrorMsg string = "An error occurred when closing %s database connection: %+v"
)

type sqlDBObserver struct {
	name     string
	instance *sql.DB
}

func (o sqlDBObserver) Close() {
	logging.Info(observerWaitingMsg, o.name)
	if observer.WaitRunningTimeout() {
		logging.Warn(observerForcingMsg, o.name)
	}

	logging.Info(observerClosingMsg, o.name)
	if o.instance == nil {
		logging.Error(dbNotInitializedErrorMsg)
		return
	}
	if err := o.instance.Close(); err != nil {
		logging.Error(observerClosingErrorMsg, o.name, err)
	}
}
