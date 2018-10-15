package logger

import "sync"

// Config comment
type Config struct {
	mu    *sync.Mutex
	Debug bool
}

// Logger comment
type Logger struct {
	PackageName string
	ServiceName string
	Conf        Config
}
