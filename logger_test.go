package logger

import (
	"testing"
)

func TestGetNewLogger(t *testing.T) {
	logger := New("logger", "main")

	if logger.PackageName != "LOGGER" {
		t.Errorf("Expected 'LOGGER' but got '%s'", logger.PackageName)
	}

	if logger.ServiceName != "MAIN" {
		t.Errorf("Expected 'MAIN' but got '%s'", logger.ServiceName)
	}
}
