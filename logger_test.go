package logger

import (
	"errors"
	"testing"
)

func TestGetNewLogger(t *testing.T) {
	logger := New("logger", "main")

	if logger.PackageName != "logger" {
		t.Errorf("Expected 'logger' but got '%s'", logger.PackageName)
	}

	if logger.ServiceName != "main" {
		t.Errorf("Expected 'main' but got '%s'", logger.ServiceName)
	}
}
func TestInfo(t *testing.T) {
	logger := New("logger", "main")
	logger.Info("Test Info message")
}
func TestWarn(t *testing.T) {
	logger := New("logger", "main")
	logger.Warn("Test Warn message")
}
func TestError(t *testing.T) {
	err := errors.New("This is an error")
	logger := New("logger", "main")

	logger.Error("Test Error message", err)
}
