package logger

import (
	"bytes"
	"context"
	"testing"
)

func TestNew(t *testing.T) {
	logger := New()
	if logger == nil {
		t.Fatal("Expected logger, got nil")
	}
}

func TestLogger_Debug(t *testing.T) {
	var buf bytes.Buffer
	logger := New(WithOutput(&buf), WithLevel(DebugLevel))

	logger.Debug("debug message", String("key", "value"))

	output := buf.String()
	if output == "" {
		t.Fatal("Expected output, got empty string")
	}
	if !contains(output, "debug message") {
		t.Fatalf("Expected 'debug message' in output, got: %s", output)
	}
}

func TestLogger_Info(t *testing.T) {
	var buf bytes.Buffer
	logger := New(WithOutput(&buf), WithLevel(InfoLevel))

	logger.Info("info message", String("key", "value"))

	output := buf.String()
	if output == "" {
		t.Fatal("Expected output, got empty string")
	}
	if !contains(output, "info message") {
		t.Fatalf("Expected 'info message' in output, got: %s", output)
	}
}

func TestLogger_Warn(t *testing.T) {
	var buf bytes.Buffer
	logger := New(WithOutput(&buf), WithLevel(WarnLevel))

	logger.Warn("warn message", String("key", "value"))

	output := buf.String()
	if output == "" {
		t.Fatal("Expected output, got empty string")
	}
	if !contains(output, "warn message") {
		t.Fatalf("Expected 'warn message' in output, got: %s", output)
	}
}

func TestLogger_Error(t *testing.T) {
	var buf bytes.Buffer
	logger := New(WithOutput(&buf), WithLevel(ErrorLevel))

	logger.Error("error message", ErrString("test error"), String("key", "value"))

	output := buf.String()
	if output == "" {
		t.Fatal("Expected output, got empty string")
	}
	if !contains(output, "error message") {
		t.Fatalf("Expected 'error message' in output, got: %s", output)
	}
	if !contains(output, "test error") {
		t.Fatalf("Expected 'test error' in output, got: %s", output)
	}
}

func TestLogger_WithFields(t *testing.T) {
	var buf bytes.Buffer
	logger := New(WithOutput(&buf), WithLevel(InfoLevel))

	logger = logger.WithFields(String("global", "value"))
	logger.Info("test message", String("local", "local_value"))

	output := buf.String()
	if output == "" {
		t.Fatal("Expected output, got empty string")
	}
	if !contains(output, "test message") {
		t.Fatalf("Expected 'test message' in output, got: %s", output)
	}
	if !contains(output, "global") || !contains(output, "value") {
		t.Fatalf("Expected global field in output, got: %s", output)
	}
	if !contains(output, "local") || !contains(output, "local_value") {
		t.Fatalf("Expected local field in output, got: %s", output)
	}
}

func TestLogger_WithContext(t *testing.T) {
	logger := New()
	ctx := logger.WithContext(context.Background())

	// 从context中获取logger
	retrievedLogger := FromContext(ctx)
	if retrievedLogger == nil {
		t.Fatal("Expected logger from context, got nil")
	}
}

func TestLogger_Clone(t *testing.T) {
	logger := New()
	cloned := logger.Clone()

	if cloned == nil {
		t.Fatal("Expected cloned logger, got nil")
	}
}

func TestLogger_Level(t *testing.T) {
	// 测试不同日志级别
	levels := []Level{DebugLevel, InfoLevel, WarnLevel, ErrorLevel, DPanicLevel, PanicLevel, FatalLevel}

	for _, level := range levels {
		var buf bytes.Buffer
		logger := New(WithOutput(&buf), WithLevel(level))

		// 如果当前级别高于info，info日志应该不输出
		if level > InfoLevel {
			logger.Info("should not appear")
			output := buf.String()
			if contains(output, "should not appear") {
				t.Fatalf("Expected no output for level %v, got: %s", level, output)
			}
		}

		// 如果当前级别低于error，error日志应该输出
		if level < ErrorLevel {
			buf.Reset()
			logger.Error("should appear", ErrString("test error"))
			output := buf.String()
			if !contains(output, "should appear") {
				t.Fatalf("Expected output for level %v, got: %s", level, output)
			}
		}
	}
}

func TestString(t *testing.T) {
	fields := String("key", "value")
	if fields["key"] != "value" {
		t.Fatalf("Expected 'value', got: %v", fields["key"])
	}
}

func TestInt(t *testing.T) {
	fields := Int("number", 42)
	if fields["number"] != 42 {
		t.Fatalf("Expected 42, got: %v", fields["number"])
	}
}

func TestBool(t *testing.T) {
	fields := Bool("flag", true)
	if fields["flag"] != true {
		t.Fatalf("Expected true, got: %v", fields["flag"])
	}
}

func TestError(t *testing.T) {
	err := ErrString("test error")
	fields := Error(err)
	if fields["error"] != "test error" {
		t.Fatalf("Expected 'test error', got: %v", fields["error"])
	}
}

func TestFormatFields(t *testing.T) {
	fields1 := String("key1", "value1")
	fields2 := String("key2", "value2")
	fields3 := String("key1", "override") // 覆盖key1

	result := FormatFields(fields1, fields2, fields3)

	if result["key1"] != "override" {
		t.Fatalf("Expected 'override', got: %v", result["key1"])
	}
	if result["key2"] != "value2" {
		t.Fatalf("Expected 'value2', got: %v", result["key2"])
	}
}

func TestParseFormat(t *testing.T) {
	msg, fields := ParseFormat("Hello %s, you are %d years old", "Alice", 25)

	if msg != "Hello Alice, you are 25 years old" {
		t.Fatalf("Expected 'Hello Alice, you are 25 years old', got: %s", msg)
	}

	if fields["arg1"] != "Alice" {
		t.Fatalf("Expected 'Alice' in fields, got: %v", fields["arg1"])
	}
	if fields["arg2"] != 25 {
		t.Fatalf("Expected 25 in fields, got: %v", fields["arg2"])
	}
}

func TestStartTimer(t *testing.T) {
	timer := StartTimer()
	fields := make(Fields)
	timer(fields)

	if _, ok := fields["duration"]; !ok {
		t.Fatal("Expected 'duration' in fields")
	}
}

func TestNewDevelopment(t *testing.T) {
	logger := NewDevelopment()
	if logger == nil {
		t.Fatal("Expected development logger, got nil")
	}
}

func TestNewProduction(t *testing.T) {
	logger := NewProduction()
	if logger == nil {
		t.Fatal("Expected production logger, got nil")
	}
}

// 辅助函数
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) &&
		(s[:len(substr)] == substr || s[len(s)-len(substr):] == substr ||
			findSubstring(s, substr)))
}

func findSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func ErrString(err string) error {
	return &testError{msg: err}
}

type testError struct {
	msg string
}

func (e *testError) Error() string {
	return e.msg
}
