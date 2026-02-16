package database

import (
	_ "context"
	"database/sql"
	"errors"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sdkopen/sdkopen-go/common/env"
	"github.com/sdkopen/sdkopen-go/logging"
)

const (
	migrationDefaultPath string = "./database/migrations"

	migrationIgnoringMsg           string = "Ignoring migration because env variable SQL_DB_MIGRATION is set to false"
	migrationStartingMsg           string = "Starting migration execution"
	migrationCouldNotConnectDBMsg  string = "Could not connect to database for migration: %v"
	migrationExecutionWithErrorMsg string = "An error when executing database migration: %v"
	migrationFinalizedMsg          string = "Migration finalized successfully"
)

func migration(db *sql.DB) error {
	if !env.SQL_DB_EXEC_MIGRATION {
		logging.Info(migrationIgnoringMsg)
		return nil
	}

	sourceUrl := env.SQL_DB_MIGRATION_SOURCE_URL
	if sourceUrl == "" {
		sourceUrl = migrationDefaultPath
	}

	logging.Info(migrationStartingMsg)
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		logging.Error(migrationCouldNotConnectDBMsg, err)
		return err
	}

	logging.Info(migrationStartingMsg)
	migrationDBInstance, _ := migrate.NewWithDatabaseInstance("file://"+sourceUrl, env.SQL_DB_NAME, driver)
	if migrationDBInstance != nil {
		if err = migrationDBInstance.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			logging.Error(migrationExecutionWithErrorMsg, err)
			return err
		}
	}

	logging.Info(migrationFinalizedMsg)
	return nil
}
