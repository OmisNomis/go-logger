package anotherpackage

import l "github.com/OmisNomis/go-logger"

var logger = l.Log("TestPackage")

// RunMe Comment
func RunMe() {
	logger.Debugf("This is a another package log with %s", "Different Args")
}
