package logger

import (
	"runtime/debug"
	"strings"
)

func (l *Logger) getStack() string {
	return strings.Join(strings.Split(string(debug.Stack()), "\n")[7:], "\n")
}
