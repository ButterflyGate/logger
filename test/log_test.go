package logger_test

import (
	"testing"

	"github.com/ButterflyGate/logger"
	"github.com/ButterflyGate/logger/levels"
)

func TestOptionLogger(t *testing.T) {
	o := logger.DefaultOutputOption() //.HideCursor().HideLevel().HideTimestamp()
	of := logger.DefaultFormatOption().FormatOneline()
	l := logger.NewLogger(
		levels.Trace,
		o, of,
	)
	l.Trace("hello,world")
	l.Alert(l)
}
