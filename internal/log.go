package internal

import (
	"fmt"
	"os"

	"github.com/ButterflyGate/logger/internal/options"
	. "github.com/ButterflyGate/logger/levels"
)

type Logger struct {
	level   LogLevel
	options *options.Options

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
func NewLoggerWithOption(level LogLevel, option *options.Options) *Logger {
	if option == nil {
		option = options.NewDefaultOptions()
	}

	l := &Logger{
		level:         level,
		options:       option,
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
		fmt.Fprintf(os.Stderr, "unknown log level %v", level)
		return l
	}
	l.Informational("successfly created logger struct")
	return l
}

func (l *Logger) ResetLevel(level LogLevel) {
	l = NewLoggerWithOption(level, l.options)
}

// func fillSpace(wordCount int, msg string) string {
// 	if len(msg) <= wordCount {
// 		return msg
// 	}
// 	space := make([]byte, wordCount)
// 	msg = string(append([]byte(msg), space...))
// 	return msg[:wordCount]
// }

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
func (l *Logger) Informational(msg any, args ...any) {
	l.informational(l.options, msg, args...)
}
func (l *Logger) Debug(msg any, args ...any) {
	l.debug(l.options, msg, args...)
}
func (l *Logger) Trace(msg any, args ...any) {
	l.trace(l.options, msg, args...)
}
