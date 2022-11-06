package logger_test

import (
	"testing"
	"time"

	"github.com/ButterflyGate/logger"
	. "github.com/ButterflyGate/logger/levels"
	"github.com/ButterflyGate/logger/options"
)

func TestLogger(t *testing.T) {
	type aa struct {
		ID       int
		FullName string
		name     string
		Birth    time.Time
	}
	a := aa{
		ID:       10,
		FullName: "kyota tahsiro",
		name:     "田代",
		Birth:    time.Date(1995, 7, 19, 0, 0, 0, 0, time.Local),
	}

	l := logger.NewLogger(
		Notice,
	)
	l.Notice(a)
}

func TestOptionLogger(t *testing.T) {
	type aa struct {
		ID       int
		FullName string
		name     string
		Birth    time.Time
	}
	a := aa{
		ID:       10,
		FullName: "kyota tahsiro",
		name:     "田代",
		Birth:    time.Date(1995, 7, 19, 0, 0, 0, 0, time.Local),
	}

	l := logger.NewLoggerWithOption(
		Notice,
		options.NewOutputOptions(false, false, false, true, true, true),
	)
	l.Notice(a)
}

func TestOptionLogger2(t *testing.T) {
	type aa struct {
		ID       int
		FullName string
		name     string
		Birth    time.Time
	}
	a := aa{
		ID:       10,
		FullName: "kyota tahsiro",
		name:     "田代",
		Birth:    time.Date(1995, 7, 19, 0, 0, 0, 0, time.Local),
	}

	l := logger.NewLoggerWithOption(
		Notice,
		&options.OutputHideOptions{},
	)
	l.Notice(a)
}
