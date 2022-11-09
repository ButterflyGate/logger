package options

type Options interface {
	IsOutputLevel() bool
	IsOutputCursor() bool
	IsOutputTimestamp() bool

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

func NewController(options ...Child) Options {
	c := &controller{
		output: NewOutputOption(),
		format: NewFormatOption(),
	}

	for i := range options {
		options[i].apply(c)
	}
	return c
}

func (o *controller) IsOutputLevel() bool {
	return o.output.level
}

func (o *controller) IsOutputCursor() bool {
	return o.output.cursor
}
func (o *controller) IsOutputTimestamp() bool {
	return o.output.timestamp
}
func (o *controller) IsFormatJson() bool {
	return o.format.json
}

func (o *controller) IsFormatReadable() bool {
	return o.format.readable
}
