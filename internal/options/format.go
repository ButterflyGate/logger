package options

type FormatOptions interface {
	FormatJson() FormatOptions
	FormatText() FormatOptions

	apply(*controller) *controller
}

type formatOptions struct {
	json     bool
	readable bool
}

func NewFormatOption() formatOptions {
	return formatOptions{
		json:     true,
		readable: true,
	}
}

func (o formatOptions) FormatJson() FormatOptions {
	o.json = true
	return o
}

func (o formatOptions) FormatText() FormatOptions {
	o.json = false
	return o
}

func (o formatOptions) FormatReadable() FormatOptions {
	o.json = false
	return o
}

func (o formatOptions) FormatOneline() FormatOptions {
	o.json = false
	return o
}

func (o formatOptions) apply(parent *controller) *controller {
	if parent == nil {
		parent = &controller{}
	}
	parent.format = o
	return parent
}
