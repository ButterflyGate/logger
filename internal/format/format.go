package format

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/ButterflyGate/logger/internal/options"
)

// type LogFormat interface{
// 	String()string
// }

type LogFormat struct {
	ctrl       options.Controller `json:"-"`
	Level      string             `json:"level,omitempty"`
	Timestamp  *time.Time         `json:"timestamp,omitempty"`
	Cursor     string             `json:"cursor,omitempty"`
	Message    interface{}        `json:"message,omitempty"`
	Data       interface{}        `json:"structure_data,omitempty"`
	StructName string             `json:"structure_name,omitempty"`
}

func NewLogFormat(ctrl options.Controller, level string, timestamp *time.Time, mainMsg interface{}, args ...interface{}) *LogFormat {
	if ctrl == nil {
		ctrl = options.NewController()
	}
	var data interface{}
	varName := ""
	cursor := getCursor()
	switch msg := mainMsg.(type) {
	case error:
		mainMsg = errorTypeMsg(ctrl, msg, args...)
	case string:
		mainMsg = stringTypeMsg(ctrl, msg, args...)
	case map[string]interface{}:
		mainMsg, data, varName = otherTypeMsg(msg, args...)
	default:
		mainMsg, data, varName = otherTypeMsg(msg, args...)
	}

	return &LogFormat{
		ctrl:       ctrl,
		Level:      ctrl.OutputLevel(level),
		Timestamp:  ctrl.OutputTimestamp(timestamp),
		Cursor:     ctrl.OutputCursor(cursor),
		Message:    ctrl.OutputMessage(mainMsg),
		Data:       ctrl.OutputData(data),
		StructName: ctrl.OutputStructName(varName),
	}
}

func (l *LogFormat) String() string {
	format, err := []byte{}, error(nil)
	if l.ctrl.IsFormatReadable() {
		format, err = json.MarshalIndent(l, "", "  ")
	} else {
		format, err = json.Marshal(l)
	}

	if err != nil {
		panic(map[string]interface{}{
			"message": err.Error(),
			"data":    l,
		})
	}
	return string(format)
}

func stringTypeMsg(ctrl options.Controller, format string, a ...any) interface{} {
	msg := fmt.Sprintf(format, a...)
	if ctrl.IsFormatMsgRowLimitted() {
		return strings.Split(msg, "\n")[:ctrl.LimitRowNum()]
	}
	return strings.Split(msg, "\n")
}

func errorTypeMsg(ctrl options.Controller, err error, a ...any) interface{} {
	msg := fmt.Sprintf("%+v", err)
	for _, v := range a {
		msg = fmt.Sprintf("%v\n%v", msg, v)
	}
	if ctrl.IsFormatMsgRowLimitted() {
		return strings.Split(msg, "\n")[:ctrl.LimitRowNum()]
	}
	return strings.Split(msg, "\n")
}

func otherTypeMsg(data interface{}, args ...interface{}) (interface{}, interface{}, string) {
	name := "unknown"
	if len(args) >= 1 {
		name = args[0].(string)
	}
	if len(args) <= 1 {
		return nil, data, name
	}

	msg := make([]string, 0, len(args))
	for i, v := range args[1:] {
		msg = append(msg, fmt.Sprintf("unknown message argument index%02d: %v", i+3, v))
	}
	return msg, data, name
}

func getCursor() string {
	_, file, line, _ := runtime.Caller(4)
	return file + ":" + strconv.Itoa(line)
}
