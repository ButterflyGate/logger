package logger

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"runtime"

	"strconv"
	"time"

	"github.com/ButterflyGate/logger/format"
	. "github.com/ButterflyGate/logger/levels"
)

type mylog func(msg any, args ...any)

func makeLogFunc(l LogLevel, output io.Writer) mylog {
	return mylog(
		func(mainMsg any, args ...any) {

			defer func() {
				err := recover()
				if err != nil {
					recovery(err, mainMsg)
				}
			}()
			now := time.Now()

			logMsg := format.NewLogFormat(
				l, now, mainMsg, args...,
			)
			fmt.Fprintln(output, logMsg)
		},
	)
}

func noneLog(msg any, args ...any) {}

type Logger struct {
	level         LogLevel
	emergency     mylog
	alert         mylog
	critical      mylog
	errors        mylog
	warning       mylog
	notice        mylog
	informational mylog
	debug         mylog
	trace         mylog
}

func NewLogger(level LogLevel) *Logger {
	l := &Logger{
		level:         level,
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
	l.informational("successfly created logger struct")
	return l
}

// func fillSpace(wordCount int, msg string) string {
// 	if len(msg) <= wordCount {
// 		return msg
// 	}
// 	space := make([]byte, wordCount)
// 	msg = string(append([]byte(msg), space...))
// 	return msg[:wordCount]
// }

func recovery(err, msg any) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Fprintln(os.Stderr,
				`{
	"level": "FATAL",
	"fatal-message1": "FATAL",
	"fatal-message2": "Logger Could Not Output A Log Message",
	"fatal-message3": "The Log Message Is Lost",
	"fatal-message4": "Probably This System Is Recovered And Is Complete"
}`)
		}
	}()

	now := time.Now()
	_, file, line, _ := runtime.Caller(18)
	cursor := file + ":" + strconv.Itoa(line)
	fmt.Fprintf(os.Stderr,
		`{
	"level": "Fatal",
	"timestamp": %v,
	"cursor": %s,
	"fatal-message": "Logger Not Working and Trying Recovering"
	"message": %+v
}`+"\n", now, cursor, msg)

	je := json.NewEncoder(os.Stderr)
	je.Encode(err)

	now = time.Now()
	fmt.Fprintf(os.Stderr,
		`{
		"level": "Fatal",
		"timestamp": %v,
		"cursor": %s
		"fatal-message": "Success Recovering"
}`+"\n", now, cursor)
}

func (l *Logger) Emergency(msg any, args ...any) {
	l.emergency(msg, args...)
}
func (l *Logger) Alert(msg any, args ...any) {
	l.alert(msg, args...)
}
func (l *Logger) Critical(msg any, args ...any) {
	l.critical(msg, args...)
}
func (l *Logger) Error(msg any, args ...any) {
	l.errors(msg, args...)
}
func (l *Logger) Warning(msg any, args ...any) {
	l.warning(msg, args...)
}
func (l *Logger) Notice(msg any, args ...any) {
	l.notice(msg, args...)
}
func (l *Logger) Informational(msg any, args ...any) {
	l.informational(msg, args...)
}
func (l *Logger) Debug(msg any, args ...any) {
	l.debug(msg, args...)
}
func (l *Logger) Trace(msg any, args ...any) {
	l.trace(msg, args...)
}
