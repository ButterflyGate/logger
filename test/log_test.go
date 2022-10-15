package logger_test

import (
	"testing"
	"time"

	"github.com/ButterflyGate/logger"
	. "github.com/ButterflyGate/logger/levels"
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
