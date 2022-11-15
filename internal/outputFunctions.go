package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/ButterflyGate/logger/internal/format"
	"github.com/ButterflyGate/logger/internal/options"
	. "github.com/ButterflyGate/logger/levels"
)

type logFunc func(option options.Controller, msg any, args ...any)

func makeLogFunc(l LogLevel, output io.Writer) logFunc {
	level := l.String()
	return logFunc(
		func(option options.Controller, mainMsg any, args ...any) {
			defer func() {
				err := recover()
				if err != nil {
					recovery(err, mainMsg)
				}
			}()
			now := time.Now()

			logMsg := format.NewLogFormat(
				option, level, &now, mainMsg, args...,
			)

			fmt.Fprintln(output, logMsg)
		},
	)
}

func noneLog(_ options.Controller, _ any, _ ...any) {}

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
	_, file, line, _ := runtime.Caller(6)
	cursor := file + ":" + strconv.Itoa(line)
	fmt.Fprintf(os.Stderr,
		`{
	"level": "Fatal",
	"timestamp": %v,
	"cursor": %s,
	"fatal-message": "Logger Not Working and Trying Recovering"
	"message": %+v
}
`, now, cursor, msg)

	je := json.NewEncoder(os.Stderr)
	je.Encode(err)

	now = time.Now()
	fmt.Fprintf(os.Stderr,
		`{
		"level": "Fatal",
		"timestamp": %v,
		"cursor": %s
		"fatal-message": "Success Recovering"
}
`, now, cursor)
}
