/**
 * @Author: dingQingHui
 * @Description:
 * @File: config
 * @Version: 1.0.0
 * @Date: 2024/9/23 11:14
 */

package zlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
)

type Option func(*Options)

func loadOptions(options ...Option) *Options {
	opts := new(Options)
	for _, option := range options {
		option(opts)
	}
	return opts
}

type Options struct {
	path         string
	level        zapcore.Level
	printConsole bool
	zapOption    []zap.Option
	writer       io.Writer
}

func (o *Options) getWriter() io.Writer {
	if o.writer != nil {
		return o.writer
	}
	return defaultWriter(o.getPath())
}

func (o *Options) getPath() string {
	if o.path != "" {
		return o.path
	}
	return "./log"
}

func (o *Options) getLevel() zapcore.Level {
	return o.level
}
func (o *Options) getZapOption() []zap.Option {
	return o.zapOption
}
func (o *Options) getPrintConsole() bool {
	return o.printConsole
}

func WithPath(path string) Option {
	return func(op *Options) {
		op.path = path
	}
}

func WithLevel(level zapcore.Level) Option {
	return func(op *Options) {
		op.level = level
	}
}
func WithPrintConsole(printConsole bool) Option {
	return func(op *Options) {
		op.printConsole = printConsole
	}
}

func WithZapOption(zapOption []zap.Option) Option {
	return func(op *Options) {
		op.zapOption = zapOption
	}
}

func WithWrite(writer io.Writer) Option {
	return func(op *Options) {
		op.writer = writer
	}
}
