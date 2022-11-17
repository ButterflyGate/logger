package internal

import (
	"os"

	"github.com/ButterflyGate/logger/internal/options"
	"github.com/ButterflyGate/logger/levels"
	. "github.com/ButterflyGate/logger/levels"
	"golang.org/x/xerrors"
)

type Logger struct {
	level   LogLevel
	options options.Controller

	emergency logFunc
	alert     logFunc
	crit      logFunc
	err       logFunc
	warn      logFunc
	notice    logFunc
	info      logFunc
	debug     logFunc
	trace     logFunc
}

func NewLogger(level LogLevel) *Logger {
	return NewLoggerWithOption(level, nil)
}
func NewLoggerWithOption(level LogLevel, option ...options.Child) *Logger {
	c := options.NewController(option...)
	l := &Logger{
		level:   level,
		options: c,
	}
	err := l.setFunction(level)
	if err != nil {
		l.setFunction(levels.Info)
		l.Warn(err)
		l.Info("log level is changed to \"%s\"", levels.Info)
		return l
	}
	l.Info("successfly created logger struct")
	return l
}

func (l *Logger) setFunction(level LogLevel) error {
	*l = Logger{
		level:     level,
		options:   l.options,
		emergency: noneLog,
		alert:     noneLog,
		crit:      noneLog,
		err:       noneLog,
		warn:      noneLog,
		notice:    noneLog,
		info:      noneLog,
		debug:     noneLog,
		trace:     noneLog,
	}

	switch level {
	case Trace:
		l.trace = makeLogFunc(Trace, os.Stdout)
		fallthrough
	case Debug:
		l.debug = makeLogFunc(Debug, os.Stdout)
		fallthrough
	case Info:
		l.info = makeLogFunc(Info, os.Stdout)
		fallthrough
	case Notice:
		l.notice = makeLogFunc(Notice, os.Stdout)
		fallthrough
	case Warn:
		l.warn = makeLogFunc(Warn, os.Stderr)
		fallthrough
	case Error:
		l.err = makeLogFunc(Error, os.Stderr)
		fallthrough
	case Crit:
		l.crit = makeLogFunc(Crit, os.Stderr)
		fallthrough
	case Alert:
		l.alert = makeLogFunc(Alert, os.Stderr)
		fallthrough
	case Emergency:
		l.emergency = makeLogFunc(Emergency, os.Stderr)
	default:
		return xerrors.Errorf("unknown log level %v", level)
	}
	return nil
}

func (l *Logger) ResetLevel(level LogLevel) {
	err := l.setFunction(level)
	if err != nil {
		l.setFunction(levels.Info)
		l.Warn(err)
		l.Info("log level is changed to \"%s\"", levels.Info)
		return
	}
	l.Info("set loglevel %s", level)
}

func (l *Logger) Emergency(msg any, args ...any) {
	l.emergency(l.options, msg, args...)
}
func (l *Logger) Alert(msg any, args ...any) {
	l.alert(l.options, msg, args...)
}
func (l *Logger) Crit(msg any, args ...any) {
	l.crit(l.options, msg, args...)
}
func (l *Logger) Error(msg any, args ...any) {
	l.err(l.options, msg, args...)
}
func (l *Logger) Warn(msg any, args ...any) {
	l.warn(l.options, msg, args...)
}
func (l *Logger) Notice(msg any, args ...any) {
	l.notice(l.options, msg, args...)
}
func (l *Logger) Info(msg any, args ...any) {
	l.info(l.options, msg, args...)
}
func (l *Logger) Debug(msg any, args ...any) {
	l.debug(l.options, msg, args...)
}
func (l *Logger) Trace(msg any, args ...any) {
	l.trace(l.options, msg, args...)
}
