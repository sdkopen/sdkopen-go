package database

import (
	"database/sql"

	"github.com/sdkopen/sdkopen-go/common/observer"
	"github.com/sdkopen/sdkopen-go/logging"
)

const (
	dbNotInitializedErrorMsg string = "database not initialized"
)

var dbInstance *sql.DB

func Initialize(factory func() *sql.DB) {
	dbInstance = factory()

	if err := observer.Attach(databaseObserver{dbInstance}); err != nil {
		logging.Fatal("could not attach database to observer: %v", err)
		return
	}
	logging.Info("database connected")
}
