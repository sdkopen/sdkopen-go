package database

import (
	"database/sql"
	"fmt"

	"github.com/sdkopen/sdkopen-go/common/env"
	"github.com/sdkopen/sdkopen-go/logging"

	_ "github.com/lib/pq"
)

const (
	defaultDriver        string = "postgres"
	defaultConnectionURI string = "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s"
	dbMigrationErrorMsg  string = "An error occurred when validate database migrations: %+v"
)

type PostgresqlConnector struct {
	host     string
	port     int
	database string
	username string
	password string
	sslMode  string
}

func NewDefaultPostgresqlConnector() *PostgresqlConnector {
	return &PostgresqlConnector{
		host:     env.SQL_DB_URL,
		port:     env.SQL_DB_PORT,
		database: env.SQL_DB_NAME,
		username: env.SQL_DB_USERNAME,
		password: env.SQL_DB_PASSWORD,
		sslMode:  env.SQL_DB_SSL_MODE,
	}
}

func (c *PostgresqlConnector) Connect() *sql.DB {
	var connectionString = c.getConnectionURI()

	db, err := sql.Open(env.SQL_DB_DRIVER, connectionString)
	if err != nil {
		logging.Fatal("an error occurred while trying to connect to the %s database: %+v", defaultDriver, err)
	}

	if err = db.Ping(); err != nil {
		logging.Fatal("an error occurred while trying to connect to the %s database: %+v", defaultDriver, err)
	}

	if err := migration(db); err != nil {
		logging.Fatal(dbMigrationErrorMsg, err)
	}

	return db
}

func (c *PostgresqlConnector) getConnectionURI() string {
	return fmt.Sprintf(defaultConnectionURI,
		c.host,
		c.port,
		c.username,
		c.password,
		c.database,
		c.sslMode)
}

func Postgresql() *sql.DB {
	return NewDefaultPostgresqlConnector().Connect()
}
