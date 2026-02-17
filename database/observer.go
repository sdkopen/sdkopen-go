package database

import (
	"database/sql"

	"github.com/sdkopen/sdkopen-go/common/observer"
	"github.com/sdkopen/sdkopen-go/logging"
)

type databaseObserver struct {
	instance *sql.DB
}

func (o databaseObserver) Close() {
	logging.Info("waiting to safely close the database connection")
	if observer.WaitRunningTimeout() {
		logging.Warn("WaitGroup timed out, forcing close the database connection")
	}

	logging.Info("closing database connection")
	if o.instance == nil {
		logging.Error(dbNotInitializedErrorMsg)
		return
	}
	if err := o.instance.Close(); err != nil {
		logging.Error("an error occurred when closing database connection: %+v", err)
	}
}
