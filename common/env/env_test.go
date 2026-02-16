package env

import (
	"os"
	"testing"
)

func TestConvertToInt_ValidValue(t *testing.T) {
	t.Setenv("TEST_INT_PORT", "3000")

	var result int
	err := convertToInt(&result, "TEST_INT_PORT")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result != 3000 {
		t.Fatalf("expected 3000, got %d", result)
	}
}

func TestConvertToInt_EmptyValue(t *testing.T) {
	os.Unsetenv("TEST_INT_EMPTY")

	var result int = 42
	err := convertToInt(&result, "TEST_INT_EMPTY")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result != 42 {
		t.Fatalf("expected 42 (unchanged), got %d", result)
	}
}

func TestConvertToInt_InvalidValue(t *testing.T) {
	t.Setenv("TEST_INT_INVALID", "not_a_number")

	var result int
	err := convertToInt(&result, "TEST_INT_INVALID")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestConvertBoolEnv_TrueValue(t *testing.T) {
	t.Setenv("TEST_BOOL_TRUE", "true")

	var result bool
	err := convertBoolEnv(&result, "TEST_BOOL_TRUE")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !result {
		t.Fatal("expected true, got false")
	}
}

func TestConvertBoolEnv_FalseValue(t *testing.T) {
	t.Setenv("TEST_BOOL_FALSE", "false")

	var result bool = true
	err := convertBoolEnv(&result, "TEST_BOOL_FALSE")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result {
		t.Fatal("expected false, got true")
	}
}

func TestConvertBoolEnv_EmptyValue(t *testing.T) {
	os.Unsetenv("TEST_BOOL_EMPTY")

	var result bool = true
	err := convertBoolEnv(&result, "TEST_BOOL_EMPTY")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !result {
		t.Fatal("expected true (unchanged), got false")
	}
}

func TestConvertBoolEnv_InvalidValue(t *testing.T) {
	t.Setenv("TEST_BOOL_INVALID", "not_a_bool")

	var result bool
	err := convertBoolEnv(&result, "TEST_BOOL_INVALID")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestValidateAndLoad_ValidEnvVars(t *testing.T) {
	t.Setenv("SERVER_PORT", "9090")
	t.Setenv("SQL_DB_PORT", "5433")
	t.Setenv("RABBITMQ_PORT", "5673")
	t.Setenv("SQL_DB_EXEC_MIGRATION", "true")

	err := validateAndLoad()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if SERVER_PORT != 9090 {
		t.Fatalf("expected SERVER_PORT=9090, got %d", SERVER_PORT)
	}
	if SQL_DB_PORT != 5433 {
		t.Fatalf("expected SQL_DB_PORT=5433, got %d", SQL_DB_PORT)
	}
	if RABBITMQ_PORT != 5673 {
		t.Fatalf("expected RABBITMQ_PORT=5673, got %d", RABBITMQ_PORT)
	}
	if !SQL_DB_EXEC_MIGRATION {
		t.Fatal("expected SQL_DB_EXEC_MIGRATION=true, got false")
	}
}

func TestValidateAndLoad_InvalidServerPort(t *testing.T) {
	t.Setenv("SERVER_PORT", "invalid")

	err := validateAndLoad()
	if err == nil {
		t.Fatal("expected error for invalid SERVER_PORT, got nil")
	}
}

func TestValidateAndLoad_InvalidSQLDBPort(t *testing.T) {
	t.Setenv("SERVER_PORT", "8080")
	t.Setenv("SQL_DB_PORT", "invalid")

	err := validateAndLoad()
	if err == nil {
		t.Fatal("expected error for invalid SQL_DB_PORT, got nil")
	}
}

func TestValidateAndLoad_InvalidRabbitMQPort(t *testing.T) {
	t.Setenv("SERVER_PORT", "8080")
	t.Setenv("SQL_DB_PORT", "5432")
	t.Setenv("RABBITMQ_PORT", "invalid")

	err := validateAndLoad()
	if err == nil {
		t.Fatal("expected error for invalid RABBITMQ_PORT, got nil")
	}
}

func TestValidateAndLoad_InvalidMigrationBool(t *testing.T) {
	t.Setenv("SERVER_PORT", "8080")
	t.Setenv("SQL_DB_PORT", "5432")
	t.Setenv("RABBITMQ_PORT", "5672")
	t.Setenv("SQL_DB_EXEC_MIGRATION", "invalid")

	err := validateAndLoad()
	if err == nil {
		t.Fatal("expected error for invalid SQL_DB_EXEC_MIGRATION, got nil")
	}
}

func TestLoad_SetsStringEnvVars(t *testing.T) {
	t.Setenv("SERVER_PORT", "8080")
	t.Setenv("SQL_DB_PORT", "5432")
	t.Setenv("RABBITMQ_PORT", "5672")
	t.Setenv("SQL_DB_EXEC_MIGRATION", "false")
	t.Setenv("SQL_DB_NAME", "testdb")
	t.Setenv("SQL_DB_SSL_MODE", "disable")
	t.Setenv("SQL_DB_URL", "localhost")
	t.Setenv("SQL_DB_USERNAME", "user")
	t.Setenv("SQL_DB_PASSWORD", "pass")
	t.Setenv("SQL_DB_DRIVER", "postgres")
	t.Setenv("RABBITMQ_URL", "rabbit-host")
	t.Setenv("RABBITMQ_USERNAME", "rabbit-user")
	t.Setenv("RABBITMQ_PASSWORD", "rabbit-pass")
	t.Setenv("RABBITMQ_VHOST", "/test")

	Load()

	if SQL_DB_NAME != "testdb" {
		t.Fatalf("expected SQL_DB_NAME=testdb, got %s", SQL_DB_NAME)
	}
	if SQL_DB_SSL_MODE != "disable" {
		t.Fatalf("expected SQL_DB_SSL_MODE=disable, got %s", SQL_DB_SSL_MODE)
	}
	if SQL_DB_URL != "localhost" {
		t.Fatalf("expected SQL_DB_URL=localhost, got %s", SQL_DB_URL)
	}
	if SQL_DB_USERNAME != "user" {
		t.Fatalf("expected SQL_DB_USERNAME=user, got %s", SQL_DB_USERNAME)
	}
	if SQL_DB_PASSWORD != "pass" {
		t.Fatalf("expected SQL_DB_PASSWORD=pass, got %s", SQL_DB_PASSWORD)
	}
	if SQL_DB_DRIVER != "postgres" {
		t.Fatalf("expected SQL_DB_DRIVER=postgres, got %s", SQL_DB_DRIVER)
	}
	if RABBITMQ_URL != "rabbit-host" {
		t.Fatalf("expected RABBITMQ_URL=rabbit-host, got %s", RABBITMQ_URL)
	}
	if RABBITMQ_USERNAME != "rabbit-user" {
		t.Fatalf("expected RABBITMQ_USERNAME=rabbit-user, got %s", RABBITMQ_USERNAME)
	}
	if RABBITMQ_PASSWORD != "rabbit-pass" {
		t.Fatalf("expected RABBITMQ_PASSWORD=rabbit-pass, got %s", RABBITMQ_PASSWORD)
	}
	if RABBITMQ_VHOST != "/test" {
		t.Fatalf("expected RABBITMQ_VHOST=/test, got %s", RABBITMQ_VHOST)
	}
}
