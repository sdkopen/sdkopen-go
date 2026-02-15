package sqlDB

import (
	"database/sql"

	"github.com/sdkopen/sdkopen-go/logging"
	"github.com/sdkopen/sdkopen-go/observer"
)

const (
	dbNotInitializedErrorMsg string = "database not initialized"
)

var sqlDBInstance *sql.DB

func Initialize() {
	sqlDBInstance = NewDefaultConnector().Connect()

	if err := observer.Attach(sqlDBObserver{defaultDriver, sqlDBInstance}); err != nil {
		logging.Fatal("could not attach sqlDB to observer: %v", err)
		return
	}
	logging.Info(dbConnectionSuccessMsg, defaultDriver)
}
