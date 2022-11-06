package options

type Options struct {
	*HideOption
	*FormatOptions
}

func NewDefaultOptions() *Options {
	return &Options{
		HideOption:    NewHideOption(),
		FormatOptions: NewFormatOption(),
	}
}
