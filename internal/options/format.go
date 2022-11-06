package options

type FormatOptions struct {
	JsonReadable bool
}

func NewFormatOption() *FormatOptions {
	return &FormatOptions{
		JsonReadable: true,
	}
}

// 下記関数群に関する注釈
// options.go に定義された Options構造体をレシーバに持ち、返り値も同様のものであり、
// 本来 options.go で定義すべき関数群だが、本質的には FormatOptions を操作する関数であるため、
// 本ファイルにて定義した。

func (o *Options) FormatJsonOneLine() *Options {
	o.JsonReadable = false
	return o
}

func (o *Options) FormatJsonReadable() *Options {
	o.JsonReadable = true
	return o
}
