package database

import (
	"testing"

	"github.com/sdkopen/sdkopen-go/common/env"
)

func TestNewDefaultPostgresqlConnector(t *testing.T) {
	env.SQL_DB_URL = "localhost"
	env.SQL_DB_PORT = 5432
	env.SQL_DB_NAME = "testdb"
	env.SQL_DB_USERNAME = "user"
	env.SQL_DB_PASSWORD = "pass"
	env.SQL_DB_SSL_MODE = "disable"

	connector := NewDefaultPostgresqlConnector()

	if connector.host != "localhost" {
		t.Fatalf("expected host=localhost, got %s", connector.host)
	}
	if connector.port != 5432 {
		t.Fatalf("expected port=5432, got %d", connector.port)
	}
	if connector.database != "testdb" {
		t.Fatalf("expected database=testdb, got %s", connector.database)
	}
	if connector.username != "user" {
		t.Fatalf("expected username=user, got %s", connector.username)
	}
	if connector.password != "pass" {
		t.Fatalf("expected password=pass, got %s", connector.password)
	}
	if connector.sslMode != "disable" {
		t.Fatalf("expected sslMode=disable, got %s", connector.sslMode)
	}
}

func TestPostgresqlConnector_GetConnectionURI(t *testing.T) {
	connector := &PostgresqlConnector{
		host:     "db.example.com",
		port:     5433,
		username: "admin",
		password: "secret",
		database: "mydb",
		sslMode:  "require",
	}

	uri := connector.getConnectionURI()
	expected := "host=db.example.com port=5433 user=admin password=secret dbname=mydb sslmode=require"

	if uri != expected {
		t.Fatalf("expected:\n%s\ngot:\n%s", expected, uri)
	}
}

func TestPostgresqlConnector_GetConnectionURI_DefaultValues(t *testing.T) {
	connector := &PostgresqlConnector{
		host:     "localhost",
		port:     5432,
		username: "postgres",
		password: "",
		database: "test",
		sslMode:  "disable",
	}

	uri := connector.getConnectionURI()
	expected := "host=localhost port=5432 user=postgres password= dbname=test sslmode=disable"

	if uri != expected {
		t.Fatalf("expected:\n%s\ngot:\n%s", expected, uri)
	}
}
