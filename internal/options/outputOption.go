package options

type OutputOption interface {
	HideLevel() OutputOption
	ShowLevel() OutputOption
	HideTimestamp() OutputOption
	ShowTimestamp() OutputOption
	HideCursor() OutputOption
	ShowCursor() OutputOption

	HideMessage() OutputOption
	ShowMessage() OutputOption
	HideData() OutputOption
	ShowData() OutputOption
	HideStructName() OutputOption
	ShowStructName() OutputOption

	apply(*controller) *controller
}

type outputOption struct {
	level      bool
	timestamp  bool
	cursor     bool
	message    bool
	data       bool
	structName bool
}

// NewHideOptions return OutputHideOptions set with all NOT hide
func NewOutputOption() outputOption {
	return outputOption{
		level:      true,
		timestamp:  true,
		cursor:     true,
		message:    true,
		data:       true,
		structName: true,
	}
}

func (o outputOption) HideLevel() OutputOption {
	o.level = false
	return o
}

func (o outputOption) ShowLevel() OutputOption {
	o.level = true
	return o
}

func (o outputOption) HideTimestamp() OutputOption {
	o.timestamp = false
	return o
}

func (o outputOption) ShowTimestamp() OutputOption {
	o.timestamp = true
	return o
}

func (o outputOption) HideCursor() OutputOption {
	o.cursor = false
	return o
}

func (o outputOption) ShowCursor() OutputOption {
	o.cursor = true
	return o
}

func (o outputOption) HideMessage() OutputOption {
	o.message = false
	return o
}

func (o outputOption) ShowMessage() OutputOption {
	o.message = true
	return o
}

func (o outputOption) HideData() OutputOption {
	o.data = false
	return o
}

func (o outputOption) ShowData() OutputOption {
	o.data = true
	return o
}

func (o outputOption) HideStructName() OutputOption {
	o.structName = false
	return o
}

func (o outputOption) ShowStructName() OutputOption {
	o.structName = true
	return o
}

func (o outputOption) apply(parent *controller) *controller {
	if parent == nil {
		parent = &controller{}
	}
	parent.output = o
	return parent
}
