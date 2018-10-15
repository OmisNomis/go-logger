package logger

import "log"

// Warn comment
func (l *Logger) Warn(msg string) {
	log.Printf("%s %s WARN: %s", l.PackageName, l.ServiceName, msg)
}
