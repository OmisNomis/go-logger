package logger

import (
	"log"
	"runtime/debug"
	"strings"
)

// Warn comment
func (l *Logger) Error(msg string, err interface{}) {
	log.Printf(
		"%s %s ERROR: %s\nAdditional information: %+v\nSTACK: %s",
		l.PackageName, l.ServiceName, msg, err, getStack(),
	)
}

func getStack() string {
	return strings.Join(strings.Split(string(debug.Stack()), "\n")[7:], "\n")
}
