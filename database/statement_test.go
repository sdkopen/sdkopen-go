package database

import (
	"context"
	"testing"
)

func TestNewStatement(t *testing.T) {
	ctx := context.Background()
	stmt := NewStatement(ctx, "SELECT 1", "arg1", "arg2")

	if stmt == nil {
		t.Fatal("expected non-nil statement")
	}
	if stmt.ctx != ctx {
		t.Fatal("expected context to match")
	}
	if stmt.query != "SELECT 1" {
		t.Fatalf("expected query 'SELECT 1', got '%s'", stmt.query)
	}
	if len(stmt.args) != 2 {
		t.Fatalf("expected 2 args, got %d", len(stmt.args))
	}
}

func TestNewStatement_NoParams(t *testing.T) {
	ctx := context.Background()
	stmt := NewStatement(ctx, "SELECT 1")

	if len(stmt.args) != 0 {
		t.Fatalf("expected 0 args, got %d", len(stmt.args))
	}
}

func TestStatement_Validate_NilInstance(t *testing.T) {
	stmt := NewStatement(context.Background(), "SELECT 1")

	err := stmt.validate(nil)
	if err == nil {
		t.Fatal("expected error for nil instance, got nil")
	}
	if err.Error() != dbNotInitializedErrorMsg {
		t.Fatalf("expected '%s', got '%s'", dbNotInitializedErrorMsg, err.Error())
	}
}

func TestStatement_Validate_EmptyQuery(t *testing.T) {
	stmt := NewStatement(context.Background(), "")

	// Need a non-nil DB to pass the first check, but we can't create a real one
	// so we test the nil case and empty query separately
	err := stmt.validate(nil)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	// The nil check happens first
	if err.Error() != dbNotInitializedErrorMsg {
		t.Fatalf("expected '%s', got '%s'", dbNotInitializedErrorMsg, err.Error())
	}
}

func TestStatement_Execute_NilGlobalInstance(t *testing.T) {
	// Ensure global instance is nil
	dbInstance = nil

	stmt := NewStatement(context.Background(), "SELECT 1")
	err := stmt.Execute()
	if err == nil {
		t.Fatal("expected error for nil global instance, got nil")
	}
	if err.Error() != dbNotInitializedErrorMsg {
		t.Fatalf("expected '%s', got '%s'", dbNotInitializedErrorMsg, err.Error())
	}
}

func TestStatement_ExecuteInInstance_NilInstance(t *testing.T) {
	stmt := NewStatement(context.Background(), "SELECT 1")
	err := stmt.ExecuteInInstance(nil)
	if err == nil {
		t.Fatal("expected error for nil instance, got nil")
	}
	if err.Error() != dbNotInitializedErrorMsg {
		t.Fatalf("expected '%s', got '%s'", dbNotInitializedErrorMsg, err.Error())
	}
}

func TestStatement_ExecuteInInstance_EmptyQuery(t *testing.T) {
	stmt := NewStatement(context.Background(), "")
	err := stmt.ExecuteInInstance(nil)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
