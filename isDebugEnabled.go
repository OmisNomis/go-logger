package logger

func (l *Logger) isDebugEnabled() bool {
	l.conf.mu.RLock()
	defer l.conf.mu.RUnlock()
	return l.conf.debug.enabled
}
