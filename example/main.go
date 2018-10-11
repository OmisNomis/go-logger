package main

import (
	"errors"

	"./nestedservice"

	"bitbucket.org/simon_ordish/logger"
)

func main() {
	logger := logger.New("Example", "Main")

	logger.Info("Logging Info from main")
	logger.Warn("Logging Warn from main")
	logger.Error("Logging Error from main", errors.New("This is a custom error"))

	nestedservice.Nested()

}
