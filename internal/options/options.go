package options

import "time"

type Controller interface {
	OutputController
	FormatController
}

type OutputController interface {
	OutputLevel(lvl string) string
	OutputCursor(cursor string) string
	OutputTimestamp(t *time.Time) *time.Time
	OutputMessage(msg interface{}) interface{}
	OutputData(data interface{}) interface{}
	OutputStructName(sn string) string
}

type FormatController interface {
	IsFormatReadable() bool
	IsFormatJson() bool
}

type Child interface {
	apply(*controller) *controller
}

type controller struct {
	output outputOption
	format formatOptions
}

func NewController(options ...Child) Controller {
	c := &controller{
		output: NewOutputOption(),
		format: NewFormatOption(),
	}

	for i := range options {
		options[i].apply(c)
	}
	return c
}

func (o *controller) OutputLevel(lvl string) string {
	if o.output.level {
		return lvl
	}
	return lvl
}
func (o *controller) OutputCursor(cursor string) string {
	if o.output.cursor {
		return cursor
	}
	return ""
}
func (o *controller) OutputTimestamp(t *time.Time) *time.Time {
	if o.output.timestamp {
		return t
	}
	return nil
}
func (o *controller) OutputMessage(msg interface{}) interface{} {
	if o.output.message {
		return msg
	}
	return ""
}
func (o *controller) OutputData(data interface{}) interface{} {
	if o.output.data {
		return data
	}
	return ""
}
func (o *controller) OutputStructName(sn string) string {
	if o.output.structName {
		return sn
	}
	return ""
}
func (o *controller) IsFormatJson() bool {
	return o.format.json
}
func (o *controller) IsFormatReadable() bool {
	return o.format.readable
}
