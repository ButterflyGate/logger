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

type LogFormat struct {
	options    options.Options `json:"-"`
	Level      string          `json:"level,omitempty"`
	Timestamp  *time.Time      `json:"timestamp,omitempty"`
	Cursor     string          `json:"cursor,omitempty"`
	Message    interface{}     `json:"message,omitempty"`
	Data       interface{}     `json:"structure_data,omitempty"`
	StructName string          `json:"structure_name,omitempty"`
}

func NewLogFormat(option options.Options, level string, timestamp *time.Time, mainMsg interface{}, args ...interface{}) *LogFormat {
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
	if option == nil {
		return &LogFormat{
			options:    option,
			Level:      level,
			Timestamp:  timestamp,
			Cursor:     getCursor(),
			Message:    mainMsg,
			Data:       data,
			StructName: varName,
		}
	}

	if !option.IsOutputLevel() {
		level = ""
	}
	if !option.IsOutputTimestamp() {
		timestamp = nil
	}
	cursor := ""
	if option.IsOutputCursor() { // 関数を呼ぶ手間を省きたいためここだけ NOT演算子なし
		cursor = getCursor()
	}

	return &LogFormat{
		options:    option,
		Level:      level,
		Timestamp:  timestamp,
		Cursor:     cursor,
		Message:    mainMsg,
		Data:       data,
		StructName: varName,
	}
}

func (l *LogFormat) String() string {
	format, err := []byte{}, error(nil)
	if l.options.IsFormatReadable() {
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
