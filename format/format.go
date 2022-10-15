package format

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/ButterflyGate/logger/levels"
)

type logFormat struct {
	Level      string      `json:"level"`
	Timestamp  time.Time   `json:"timestamp"`
	Cursor     string      `json:"cursor"`
	Message    interface{} `json:"message,omitempty"`
	Data       interface{} `json:"structure_data,omitempty"`
	StructName string      `json:"structure_name,omitempty"`
}

func NewLogFormat(level levels.LogLevel, timestamp time.Time, mainMsg interface{}, args ...interface{}) *logFormat {
	var data interface{}
	varName := ""
	switch msg := mainMsg.(type) {
	case error:
		mainMsg = errorTypeMsg(msg)
	case string:
		mainMsg = stringTypeMsg(msg, args...)
	case map[string]interface{}:
		mainMsg, data, varName = otherTypeMsg(msg, args...)
	default:
		mainMsg, data, varName = otherTypeMsg(msg, args...)
	}

	cursor := getCursor()

	return &logFormat{
		Level:      level.String(),
		Timestamp:  timestamp,
		Cursor:     cursor,
		Message:    mainMsg,
		StructName: varName,
		Data:       data,
	}
}

func (l *logFormat) String() string {
	format, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		panic(map[string]interface{}{
			"message": err.Error(),
			"data":    l,
		})
	}
	return string(format)
}

func stringTypeMsg(format string, a ...any) interface{} {
	msg := fmt.Sprintf(format, a...)
	return strings.Split(msg, "\n")
}

func errorTypeMsg(err error) interface{} {
	msg := fmt.Sprintf("%+v", err)
	return strings.Split(msg, "\n")
}

func otherTypeMsg(data interface{}, args ...interface{}) (interface{}, interface{}, string) {
	name := "unknown"
	var msg interface{} = nil
	for i, v := range args {
		if i == 0 {
			name = v.(string)
			continue
		}
		if i == 1 {
			msg = make([]string, 0, len(args))
		}
		msg = append(msg.([]string), fmt.Sprintf("unknown message argument index%2d: %v", i, v))
	}
	return msg, data, name
}

func getCursor() string {
	_, file, line, _ := runtime.Caller(4)
	return file + ":" + strconv.Itoa(line)
}
