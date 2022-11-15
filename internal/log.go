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

	emergency     logFunc
	alert         logFunc
	critical      logFunc
	errors        logFunc
	warning       logFunc
	notice        logFunc
	informational logFunc
	debug         logFunc
	trace         logFunc
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
		l.setFunction(levels.Informational)
		l.Warning(err)
		l.Information("log level is changed to \"%s\"", levels.Informational)
		return l
	}
	l.Information("successfly created logger struct")
	return l
}

func (l *Logger) setFunction(level LogLevel) error {
	*l = Logger{
		level:         level,
		options:       l.options,
		emergency:     noneLog,
		alert:         noneLog,
		critical:      noneLog,
		errors:        noneLog,
		warning:       noneLog,
		notice:        noneLog,
		informational: noneLog,
		debug:         noneLog,
		trace:         noneLog,
	}

	switch level {
	case Trace:
		l.trace = makeLogFunc(Trace, os.Stdout)
		fallthrough
	case Debug:
		l.debug = makeLogFunc(Debug, os.Stdout)
		fallthrough
	case Informational:
		l.informational = makeLogFunc(Informational, os.Stdout)
		fallthrough
	case Notice:
		l.notice = makeLogFunc(Notice, os.Stdout)
		fallthrough
	case Warning:
		l.warning = makeLogFunc(Warning, os.Stderr)
		fallthrough
	case Error:
		l.errors = makeLogFunc(Error, os.Stderr)
		fallthrough
	case Critical:
		l.critical = makeLogFunc(Critical, os.Stderr)
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
		l.setFunction(levels.Informational)
		l.Warning(err)
		l.Information("log level is changed to \"%s\"", levels.Informational)
		return
	}
	l.Information("set loglevel %s", level)
}

func (l *Logger) Emergency(msg any, args ...any) {
	l.emergency(l.options, msg, args...)
}
func (l *Logger) Alert(msg any, args ...any) {
	l.alert(l.options, msg, args...)
}
func (l *Logger) Critical(msg any, args ...any) {
	l.critical(l.options, msg, args...)
}
func (l *Logger) Error(msg any, args ...any) {
	l.errors(l.options, msg, args...)
}
func (l *Logger) Warning(msg any, args ...any) {
	l.warning(l.options, msg, args...)
}
func (l *Logger) Notice(msg any, args ...any) {
	l.notice(l.options, msg, args...)
}
func (l *Logger) Information(msg any, args ...any) {
	l.informational(l.options, msg, args...)
}
func (l *Logger) Debug(msg any, args ...any) {
	l.debug(l.options, msg, args...)
}
func (l *Logger) Trace(msg any, args ...any) {
	l.trace(l.options, msg, args...)
}
