package logger

import "log"

// Debug Comment
func (l *Logger) Debug(msg string, extra ...interface{}) {
	if l.Conf.Debug {
		log.Printf(
			"%s %s DEBUG: %s\nAdditional Information: %+v",
			l.PackageName, l.ServiceName, msg, extra,
		)
	}
}
