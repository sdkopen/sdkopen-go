package commonhttp

import (
	"encoding/json"
	"testing"
)

func TestJsonEncoder_Struct(t *testing.T) {
	type testObj struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	obj := testObj{Name: "John", Age: 30}
	result, err := JsonEncoder(obj)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	var decoded testObj
	if err := json.Unmarshal(result, &decoded); err != nil {
		t.Fatalf("expected valid JSON, got unmarshal error: %v", err)
	}
	if decoded.Name != "John" || decoded.Age != 30 {
		t.Fatalf("expected {John 30}, got {%s %d}", decoded.Name, decoded.Age)
	}
}

func TestJsonEncoder_Map(t *testing.T) {
	obj := map[string]string{"key": "value"}
	result, err := JsonEncoder(obj)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	var decoded map[string]string
	if err := json.Unmarshal(result, &decoded); err != nil {
		t.Fatalf("expected valid JSON, got: %v", err)
	}
	if decoded["key"] != "value" {
		t.Fatalf("expected key=value, got key=%s", decoded["key"])
	}
}

func TestJsonEncoder_String(t *testing.T) {
	result, err := JsonEncoder("hello")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expected := `"hello"`
	if string(result) != expected {
		t.Fatalf("expected %s, got %s", expected, string(result))
	}
}

func TestStringEncoder(t *testing.T) {
	result, err := StringEncoder("hello world")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if string(result) != "hello world" {
		t.Fatalf("expected 'hello world', got '%s'", string(result))
	}
}

func TestStringEncoder_Number(t *testing.T) {
	result, err := StringEncoder(42)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if string(result) != "42" {
		t.Fatalf("expected '42', got '%s'", string(result))
	}
}

func TestEncoder_JSON(t *testing.T) {
	obj := map[string]int{"count": 5}
	result, err := Encoder(ContentTypeJSON, obj)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	var decoded map[string]int
	if err := json.Unmarshal(result, &decoded); err != nil {
		t.Fatalf("expected valid JSON, got: %v", err)
	}
	if decoded["count"] != 5 {
		t.Fatalf("expected count=5, got count=%d", decoded["count"])
	}
}

func TestEncoder_TextPlain(t *testing.T) {
	result, err := Encoder(ContentTypeTextPlain, "plain text")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if string(result) != "plain text" {
		t.Fatalf("expected 'plain text', got '%s'", string(result))
	}
}
