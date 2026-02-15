package env

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	errorIntegerParse string = "could not parse %s, permited int value, got %v: %w"
)

var (
	SQL_DB_DRIVER   string = "postgres"
	SQL_DB_PORT            = 5432
	SQL_DB_NAME            = ""
	SQL_DB_SSL_MODE        = ""
	SQL_DB_URL             = ""
	SQL_DB_USERNAME        = ""
	SQL_DB_PASSWORD        = ""
	SERVER_PORT            = 8080
)

// Load loads and validates all environment variables. It's used in app initialization.
func Load() error {
	_ = godotenv.Load()

	if err := validateAndLoad(); err != nil {
		return err
	}

	SQL_DB_NAME = os.Getenv("SQL_DB_NAME")
	SQL_DB_SSL_MODE = os.Getenv("SQL_DB_SSL_MODE")
	SQL_DB_URL = os.Getenv("SQL_DB_URL")
	SQL_DB_USERNAME = os.Getenv("SQL_DB_USERNAME")
	SQL_DB_PASSWORD = os.Getenv("SQL_DB_PASSWORD")
	SQL_DB_DRIVER = os.Getenv("SQL_DB_DRIVER")

	return nil
}

func validateAndLoad() error {
	if err := convertToInt(&SERVER_PORT, "SERVER_PORT"); err != nil {
		return err
	}

	if err := convertToInt(&SQL_DB_PORT, "SQL_DB_PORT"); err != nil {
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
