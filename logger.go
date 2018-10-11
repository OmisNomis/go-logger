package logger

import (
	"log"
	"runtime/debug"
	"strings"
)

// Logger comment
type Logger struct {
	PackageName string
	ServiceName string
}

// New comment
func New(p string, s string) Logger {
	return Logger{
		PackageName: strings.ToUpper(p),
		ServiceName: strings.ToUpper(s),
	}
}

// Info comment
func (l *Logger) Info(msg string) {
	log.Printf("%s %s INFO: %s", l.PackageName, l.ServiceName, msg)
}

// Warn comment
func (l *Logger) Warn(msg string) {
	log.Printf("%s %s WARN: %s", l.PackageName, l.ServiceName, msg)
}

// Warn comment
func (l *Logger) Error(msg string, err interface{}) {
	log.Printf("%s %s ERROR: %s\nAdditional information: %+v\nSTACK: %s", l.PackageName, l.ServiceName, msg, err, getStack())
}

func getStack() string {
	return strings.Join(strings.Split(string(debug.Stack()), "\n")[7:], "\n")
}
