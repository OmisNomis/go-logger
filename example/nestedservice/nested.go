package nestedservice

import (
	"errors"

	"bitbucket.org/simon_ordish/logger"
)

// Nested comment
func Nested() {
	logger := logger.New("Example", "nestedService")

	logger.Info("Logging Info from nested")
	logger.Warn("Logging Warn from nested")
	logger.Error("Logging Error from nested", errors.New("This is a custom error"))

}
