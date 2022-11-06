package logger

import "github.com/ButterflyGate/logger/internal/options"

type LoggerOptions interface {
	HideLevel() LoggerOptions
	ShowLevel() LoggerOptions
	HideTimestamp() LoggerOptions
	ShowTimestamp() LoggerOptions
	HideCursor() LoggerOptions
	ShowCursor() LoggerOptions
}

type Options = options.Options

func DefaultOption() *Options {
	return options.NewDefaultOptions()
}
