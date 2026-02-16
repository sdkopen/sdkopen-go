package validator

import (
	"testing"

	"github.com/google/uuid"
)

func setupValidator() {
	Initialize()
}

func TestInitialize(t *testing.T) {
	Initialize()

	if instance == nil {
		t.Fatal("expected instance to be initialized")
	}
	if instance.validator == nil {
		t.Fatal("expected validator to be initialized")
	}
	if instance.formEncoder == nil {
		t.Fatal("expected formEncoder to be initialized")
	}
	if instance.formDecoder == nil {
		t.Fatal("expected formDecoder to be initialized")
	}
}

type testStruct struct {
	Name  string `validate:"required"`
	Email string `validate:"required,email"`
	Age   int    `validate:"gte=0,lte=150"`
}

func TestStruct_Valid(t *testing.T) {
	setupValidator()

	s := testStruct{
		Name:  "John",
		Email: "john@example.com",
		Age:   30,
	}

	err := Struct(s)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestStruct_MissingRequired(t *testing.T) {
	setupValidator()

	s := testStruct{
		Email: "john@example.com",
		Age:   30,
	}

	err := Struct(s)
	if err == nil {
		t.Fatal("expected error for missing Name, got nil")
	}
}

func TestStruct_InvalidEmail(t *testing.T) {
	setupValidator()

	s := testStruct{
		Name:  "John",
		Email: "invalid-email",
		Age:   30,
	}

	err := Struct(s)
	if err == nil {
		t.Fatal("expected error for invalid email, got nil")
	}
}

func TestStruct_InvalidAge(t *testing.T) {
	setupValidator()

	s := testStruct{
		Name:  "John",
		Email: "john@example.com",
		Age:   -1,
	}

	err := Struct(s)
	if err == nil {
		t.Fatal("expected error for negative age, got nil")
	}
}

type formStruct struct {
	Name string `form:"name"`
	Page int    `form:"page"`
}

func TestFormEncode(t *testing.T) {
	setupValidator()

	s := formStruct{
		Name: "John",
		Page: 1,
	}

	values, err := FormEncode(s)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if values["name"][0] != "John" {
		t.Fatalf("expected name=John, got %s", values["name"][0])
	}
	if values["page"][0] != "1" {
		t.Fatalf("expected page=1, got %s", values["page"][0])
	}
}

func TestFormEncodeToQueryString(t *testing.T) {
	setupValidator()

	s := formStruct{
		Name: "John",
		Page: 1,
	}

	qs, err := FormEncodeToQueryString(s)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if qs == "" {
		t.Fatal("expected non-empty query string")
	}
}

func TestFormDecode(t *testing.T) {
	setupValidator()

	values := map[string][]string{
		"name": {"Jane"},
		"page": {"5"},
	}

	var s formStruct
	err := FormDecode(&s, values)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if s.Name != "Jane" {
		t.Fatalf("expected Name=Jane, got %s", s.Name)
	}
	if s.Page != 5 {
		t.Fatalf("expected Page=5, got %d", s.Page)
	}
}

type uuidStruct struct {
	ID uuid.UUID `form:"id"`
}

func TestFormDecode_UUID(t *testing.T) {
	setupValidator()

	testUUID := "550e8400-e29b-41d4-a716-446655440000"
	values := map[string][]string{
		"id": {testUUID},
	}

	var s uuidStruct
	err := FormDecode(&s, values)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if s.ID.String() != testUUID {
		t.Fatalf("expected ID=%s, got %s", testUUID, s.ID.String())
	}
}

func TestFormDecode_InvalidUUID(t *testing.T) {
	setupValidator()

	values := map[string][]string{
		"id": {"not-a-uuid"},
	}

	var s uuidStruct
	err := FormDecode(&s, values)
	if err == nil {
		t.Fatal("expected error for invalid UUID, got nil")
	}
}
