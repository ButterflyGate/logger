package logger

import "github.com/ButterflyGate/logger/internal/options"

type option = options.Child
type outputOption = options.OutputOption
type formatOption = options.FormatOptions

func DefaultOutputOption() outputOption {
	return options.NewOutputOption()
}

func DefaultFormatOption() formatOption {
	return options.NewFormatOption()
}
