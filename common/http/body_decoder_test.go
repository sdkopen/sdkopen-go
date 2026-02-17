package commonhttp

import (
	"encoding/json"
	"testing"
)

func TestJsonDecoder_Struct(t *testing.T) {
	type testObj struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	data, _ := json.Marshal(testObj{Name: "Jane", Age: 25})
	result, err := JsonDecoder[testObj](data)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result.Name != "Jane" {
		t.Fatalf("expected Name=Jane, got %s", result.Name)
	}
	if result.Age != 25 {
		t.Fatalf("expected Age=25, got %d", result.Age)
	}
}

func TestJsonDecoder_Map(t *testing.T) {
	data := []byte(`{"key":"value","num":123}`)
	result, err := JsonDecoder[map[string]any](data)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if (*result)["key"] != "value" {
		t.Fatalf("expected key=value, got %v", (*result)["key"])
	}
}

func TestJsonDecoder_InvalidJSON(t *testing.T) {
	data := []byte(`not valid json`)
	_, err := JsonDecoder[map[string]any](data)
	if err == nil {
		t.Fatal("expected error for invalid JSON, got nil")
	}
}

func TestDecoder_JSON(t *testing.T) {
	type testObj struct {
		ID int `json:"id"`
	}

	data, _ := json.Marshal(testObj{ID: 42})
	config := &DecoderConfig{
		ContentType: ContentTypeJSON,
		Data:        data,
	}

	result, err := Decoder[testObj](config)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result.ID != 42 {
		t.Fatalf("expected ID=42, got %d", result.ID)
	}
}

func TestDecoder_NilConfig(t *testing.T) {
	result, err := Decoder[map[string]any](nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result != nil {
		t.Fatal("expected nil result for nil config")
	}
}

func TestDecoder_UnsupportedContentType(t *testing.T) {
	config := &DecoderConfig{
		ContentType: ContentTypePDF,
		Data:        []byte("some data"),
	}

	result, err := Decoder[map[string]any](config)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result != nil {
		t.Fatal("expected nil result for unsupported content type")
	}
}
