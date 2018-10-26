package subpackage

import (
	l "github.com/OmisNomis/go-logger"
)

var logger = l.Log("TestPackage")

// RunMe Comment
func RunMe() {
	logger.Warnf("This is a Subpackage log with %s", "Args")
}
