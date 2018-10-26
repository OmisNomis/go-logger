package logger

import (
	"fmt"
	"log"

	"github.com/mgutz/ansi"
)

func (l *Logger) sendToTrace(format string, msg string, args ...interface{}) {
	l.conf.mu.Lock()
	defer l.conf.mu.Unlock()

	for s := range l.conf.trace.sockets {
		if l.isRegexMatch(msg, args...) {
			_, e := s.Write([]byte(fmt.Sprintf(format, args...)))
			if e != nil {
				log.Println(ansi.Color(fmt.Sprintf("Writing client error: '%s'", e), "red"))
				delete(l.conf.trace.sockets, s)
			}
		}
	}
}
