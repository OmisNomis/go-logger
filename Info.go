package logger

import "log"

// Info comment
func (l *Logger) Info(msg string, extra ...interface{}) {
	log.Printf(
		"%s %s INFO: %s",
		l.PackageName, l.ServiceName, msg,
	)
}
