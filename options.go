package logger

import "github.com/ButterflyGate/logger/internal/options"

type option = options.Child
type OutputOption = options.OutputOption
type FormatOption = options.FormatOptions

func DefaultOutputOption() OutputOption {
	return options.NewOutputOption()
}

func DefaultFormatOption() FormatOption {
	return options.NewFormatOption()
}
