package env

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sdkopen/sdkopen-go/logging"
)

const (
	errorIntegerParse string = "could not parse %s, permited int value, got %v: %w"
	errorBooleanParse string = "could not parse %s, permited 'true' or 'false', got %v: %w"
)

var (
	SQL_DB_DRIVER               = "postgres"
	SQL_DB_PORT                 = 5432
	SQL_DB_NAME                 = ""
	SQL_DB_SSL_MODE             = ""
	SQL_DB_URL                  = ""
	SQL_DB_USERNAME             = ""
	SQL_DB_PASSWORD             = ""
	SQL_DB_EXEC_MIGRATION       = false
	SQL_DB_MIGRATION_SOURCE_URL = ""
	SERVER_PORT                 = 8080
	RABBITMQ_URL                = ""
	RABBITMQ_PORT               = 5672
	RABBITMQ_USERNAME           = ""
	RABBITMQ_PASSWORD           = ""
	RABBITMQ_VHOST              = ""
)

func Load() {
	_ = godotenv.Load()

	if err := validateAndLoad(); err != nil {
		logging.Fatal("%s", err.Error())
	}

	SQL_DB_NAME = os.Getenv("SQL_DB_NAME")
	SQL_DB_SSL_MODE = os.Getenv("SQL_DB_SSL_MODE")
	SQL_DB_URL = os.Getenv("SQL_DB_URL")
	SQL_DB_USERNAME = os.Getenv("SQL_DB_USERNAME")
	SQL_DB_PASSWORD = os.Getenv("SQL_DB_PASSWORD")
	SQL_DB_DRIVER = os.Getenv("SQL_DB_DRIVER")

	SQL_DB_MIGRATION_SOURCE_URL = os.Getenv("SQL_DB_MIGRATION_SOURCE_URL")

	RABBITMQ_URL = os.Getenv("RABBITMQ_URL")
	RABBITMQ_USERNAME = os.Getenv("RABBITMQ_USERNAME")
	RABBITMQ_PASSWORD = os.Getenv("RABBITMQ_PASSWORD")
	RABBITMQ_VHOST = os.Getenv("RABBITMQ_VHOST")
}

func validateAndLoad() error {
	if err := convertToInt(&SERVER_PORT, "SERVER_PORT"); err != nil {
		return err
	}

	if err := convertToInt(&SQL_DB_PORT, "SQL_DB_PORT"); err != nil {
		return err
	}

	if err := convertToInt(&RABBITMQ_PORT, "RABBITMQ_PORT"); err != nil {
		return err
	}

	if err := convertBoolEnv(&SQL_DB_EXEC_MIGRATION, "SQL_DB_EXEC_MIGRATION"); err != nil {
		return err
	}

	return nil
}

func convertToInt(env *int, envName string) error {
	if envString := os.Getenv(envName); envString != "" {
		var err error
		if *env, err = strconv.Atoi(envString); err != nil {
			return fmt.Errorf(errorIntegerParse, envName, envString, err)
		}
	}
	return nil
}

func convertBoolEnv(env *bool, envName string) error {
	if envString := os.Getenv(envName); envString != "" {
		var err error
		if *env, err = strconv.ParseBool(envString); err != nil {
			return fmt.Errorf(errorBooleanParse, envName, envString, err)
		}
	}
	return nil
}
