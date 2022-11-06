package options

type HideOption struct {
	Level     bool
	Timestamp bool
	Cursor    bool
	// Message    bool
	// Data       bool
	// Structname bool
}

// NewHideOptions return OutputHideOptions set with all NOT hide
func NewHideOption() *HideOption {
	return &HideOption{}
}

// 下記関数群に関する注釈
// options.go に定義された Options構造体をレシーバに持ち、返り値も同様のものであり、
// 本来 options.go で定義すべき関数群だが、本質的には HidenOptions を操作する関数であるため、
// 本ファイルにて定義した。

func (o *Options) HideLevel() *Options {
	o.Level = true
	return o
}

func (o *Options) ShowLevel() *Options {
	o.Level = false
	return o
}

func (o *Options) HideTimestamp() *Options {
	o.Timestamp = true
	return o
}

func (o *Options) ShowTimestamp() *Options {
	o.Timestamp = true
	return o
}

func (o *Options) HideCursor() *Options {
	o.Cursor = true
	return o
}

func (o *Options) ShowCursor() *Options {
	o.Cursor = true
	return o
}
