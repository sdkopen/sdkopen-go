package logging

import (
	"bytes"
	"log/slog"
	"strings"
	"testing"
)

func TestParseLevel_Debug(t *testing.T) {
	level := parseLevel("debug")
	if level != slog.LevelDebug {
		t.Fatalf("expected LevelDebug, got %v", level)
	}
}

func TestParseLevel_Info(t *testing.T) {
	level := parseLevel("info")
	if level != slog.LevelInfo {
		t.Fatalf("expected LevelInfo, got %v", level)
	}
}

func TestParseLevel_Warn(t *testing.T) {
	level := parseLevel("warn")
	if level != slog.LevelWarn {
		t.Fatalf("expected LevelWarn, got %v", level)
	}
}

func TestParseLevel_Error(t *testing.T) {
	level := parseLevel("error")
	if level != slog.LevelError {
		t.Fatalf("expected LevelError, got %v", level)
	}
}

func TestParseLevel_Unknown_DefaultsToInfo(t *testing.T) {
	level := parseLevel("unknown")
	if level != slog.LevelInfo {
		t.Fatalf("expected LevelInfo as default, got %v", level)
	}
}

func TestParseLevel_CaseInsensitive(t *testing.T) {
	level := parseLevel("DEBUG")
	if level != slog.LevelDebug {
		t.Fatalf("expected LevelDebug, got %v", level)
	}
}

func setupTestLogger() *bytes.Buffer {
	buf := &bytes.Buffer{}
	LogOutput = buf
	Initialize()
	return buf
}

func TestInfo_WritesMessage(t *testing.T) {
	t.Setenv("SDKOPEN_LOG_LEVEL", "info")
	buf := setupTestLogger()

	Info("hello %s", "world")

	output := buf.String()
	if !strings.Contains(output, "hello world") {
		t.Fatalf("expected output to contain 'hello world', got: %s", output)
	}
	if !strings.Contains(output, "INFO") {
		t.Fatalf("expected output to contain 'INFO', got: %s", output)
	}
}

func TestError_WritesMessage(t *testing.T) {
	t.Setenv("SDKOPEN_LOG_LEVEL", "error")
	buf := setupTestLogger()

	Error("something went %s", "wrong")

	output := buf.String()
	if !strings.Contains(output, "something went wrong") {
		t.Fatalf("expected output to contain 'something went wrong', got: %s", output)
	}
	if !strings.Contains(output, "ERROR") {
		t.Fatalf("expected output to contain 'ERROR', got: %s", output)
	}
}

func TestWarn_WritesMessage(t *testing.T) {
	t.Setenv("SDKOPEN_LOG_LEVEL", "warn")
	buf := setupTestLogger()

	Warn("warning: %d", 42)

	output := buf.String()
	if !strings.Contains(output, "warning: 42") {
		t.Fatalf("expected output to contain 'warning: 42', got: %s", output)
	}
	if !strings.Contains(output, "WARN") {
		t.Fatalf("expected output to contain 'WARN', got: %s", output)
	}
}

func TestDebug_WritesMessage_WhenLevelDebug(t *testing.T) {
	t.Setenv("SDKOPEN_LOG_LEVEL", "debug")
	buf := setupTestLogger()

	Debug("debug info: %v", true)

	output := buf.String()
	if !strings.Contains(output, "debug info: true") {
		t.Fatalf("expected output to contain 'debug info: true', got: %s", output)
	}
}

func TestDebug_Suppressed_WhenLevelInfo(t *testing.T) {
	t.Setenv("SDKOPEN_LOG_LEVEL", "info")
	buf := setupTestLogger()

	Debug("this should not appear")

	output := buf.String()
	if strings.Contains(output, "this should not appear") {
		t.Fatalf("expected debug message to be suppressed at info level, got: %s", output)
	}
}

func TestFatal_PanicsWithMessage(t *testing.T) {
	t.Setenv("SDKOPEN_LOG_LEVEL", "error")
	setupTestLogger()

	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("expected panic from Fatal, got none")
		}
		msg, ok := r.(string)
		if !ok {
			t.Fatalf("expected panic value to be a string, got %T", r)
		}
		if !strings.Contains(msg, "fatal error") {
			t.Fatalf("expected panic message to contain 'fatal error', got: %s", msg)
		}
	}()

	Fatal("fatal error: %s", "test")
}

func TestInitialize_DefaultLevel(t *testing.T) {
	t.Setenv("SDKOPEN_LOG_LEVEL", "")
	buf := setupTestLogger()

	Info("info message")
	Debug("debug message")

	output := buf.String()
	if !strings.Contains(output, "info message") {
		t.Fatal("expected info message to appear at default level")
	}
	if strings.Contains(output, "debug message") {
		t.Fatal("expected debug message to be suppressed at default info level")
	}
}

func TestInitialize_InvalidLevel_DefaultsToInfo(t *testing.T) {
	t.Setenv("SDKOPEN_LOG_LEVEL", "invalid_level")
	buf := setupTestLogger()

	Info("info visible")
	Debug("debug hidden")

	output := buf.String()
	if !strings.Contains(output, "info visible") {
		t.Fatal("expected info to be visible")
	}
	if strings.Contains(output, "debug hidden") {
		t.Fatal("expected debug to be hidden at default info level")
	}
}
