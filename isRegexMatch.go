package logger

import (
	"regexp"
)

func (l *Logger) isRegexMatch(r string, msg string) bool {
	match, _ := regexp.MatchString(r, msg)
	return match
}
