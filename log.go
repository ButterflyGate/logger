package logger

import (
	"github.com/ButterflyGate/logger/internal"
	"github.com/ButterflyGate/logger/levels"
)

type log interface {
	Emergency(msg any, args ...any)
	Alert(msg any, args ...any)
	Crit(msg any, args ...any)
	Error(msg any, args ...any)
	Warn(msg any, args ...any)
	Notice(msg any, args ...any)
	Info(msg any, args ...any)
	Debug(msg any, args ...any)
	Trace(msg any, args ...any)
}

type Logger interface {
	log
	ResetLevel(level levels.LogLevel)
}

func NewLogger[T levels.LogLevelType](level T, options ...option) Logger {
	return internal.NewLoggerWithOption(levels.LogLevel(level), options...)
}
