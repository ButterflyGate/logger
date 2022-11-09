package options

type OutputOption interface {
	HideLevel() OutputOption
	ShowLevel() OutputOption
	HideTimestamp() OutputOption
	ShowTimestamp() OutputOption
	HideCursor() OutputOption
	ShowCursor() OutputOption

	apply(*controller) *controller
}

type outputOption struct {
	level     bool
	timestamp bool
	cursor    bool
	// Message    bool
	// Data       bool
	// Structname bool
}

// NewHideOptions return OutputHideOptions set with all NOT hide
func NewOutputOption() outputOption {
	return outputOption{
		level:     true,
		timestamp: true,
		cursor:    true,
	}
}

func (o outputOption) HideLevel() OutputOption {
	o.level = true
	return o
}

func (o outputOption) ShowLevel() OutputOption {
	o.level = false
	return o
}

func (o outputOption) HideTimestamp() OutputOption {
	o.timestamp = true
	return o
}

func (o outputOption) ShowTimestamp() OutputOption {
	o.timestamp = true
	return o
}

func (o outputOption) HideCursor() OutputOption {
	o.cursor = true
	return o
}

func (o outputOption) ShowCursor() OutputOption {
	o.cursor = true
	return o
}

func (o outputOption) apply(parent *controller) *controller {
	if parent == nil {
		parent = &controller{}
	}
	parent.output = o
	return parent
}
