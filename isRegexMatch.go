package logger

import (
	"fmt"
	"regexp"
)

func (l *Logger) isRegexMatch(msg string, args ...interface{}) bool {
	match, _ := regexp.MatchString(l.conf.debug.regex, fmt.Sprintf(msg, args...))
	return match
}
