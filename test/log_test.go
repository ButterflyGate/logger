package logger_test

import (
	"testing"
	"time"

	"github.com/ButterflyGate/logger"
	"github.com/ButterflyGate/logger/levels"
)

func TestOptionLogger(t *testing.T) {
	o := logger.DefaultOutputOption().HideCursor().HideLevel().HideTimestamp()
	of := logger.DefaultFormatOption().FormatMessageRowLimit(0)
	l := logger.NewLogger(
		levels.Trace,
		o, of,
	)
	l.Info("hello,world\nhello\nhello")
}

func TestStructureLog(t *testing.T) {
	myinfo := struct {
		ID    int
		Name  string
		Birth time.Time
	}{
		ID:    100,
		Name:  "kyota",
		Birth: time.Date(1995, 7, 19, 0, 0, 0, 0, time.Local),
	}

	l := logger.NewLogger(
		levels.Trace,
	)
	l.Info(myinfo, "myinfo", "aserf", 123)
}
