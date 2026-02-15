package sqlDB

import (
	"database/sql"
	"fmt"
	_ "time"

	"github.com/lib/pq"
	"github.com/sdkopen/sdkopen-go/common/env"
	"github.com/sdkopen/sdkopen-go/logging"
	sqltrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql"
)

const (
	defaultDriver        string = "postgres"
	defaultConnectionURI string = "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s"

	dbConnectionSuccessMsg string = "%s database connected"
	dbConnectionErrorMsg   string = "An error occurred while trying to connect to the %s database. Error: %+v"
)

type SqlConnector struct {
	host     string
	port     int
	database string
	username string
	password string
	sslMode  string
}

func NewDefaultConnector() *SqlConnector {
	return &SqlConnector{
		host:     env.SQL_DB_URL,
		port:     env.SQL_DB_PORT,
		database: env.SQL_DB_NAME,
		username: env.SQL_DB_USERNAME,
		password: env.SQL_DB_PASSWORD,
		sslMode:  env.SQL_DB_SSL_MODE,
	}
}

func (c *SqlConnector) Connect() *sql.DB {
	sqltrace.Register(defaultDriver, &pq.Driver{})

	var connectionString = c.getConnectionURI()

	sqlDB, err := sqltrace.Open(env.SQL_DB_DRIVER, connectionString)
	if err != nil {
		logging.Fatal(dbConnectionErrorMsg, defaultDriver, err)
	}

	if err = sqlDB.Ping(); err != nil {
		logging.Fatal(dbConnectionErrorMsg, defaultDriver, err)
	}

	return sqlDB
}

func (c *SqlConnector) getConnectionURI() string {
	password := c.password

	return fmt.Sprintf(defaultConnectionURI,
		c.host,
		c.port,
		c.username,
		password,
		c.database,
		c.sslMode)
}
