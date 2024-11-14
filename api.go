/**
 * @Author: dingQingHui
 * @Description:
 * @File: api.go
 * @Version: 1.0.0
 * @Date: 2024/11/14 11:32
 */

package zlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type IZLogger interface {
	Debug(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Panic(msg string, fields ...zap.Field)
	Fatal(msg string, fields ...zap.Field)
	SetLogLevel(logLevel zapcore.Level)
	GetLogLevel() zapcore.Level
	Stop()
}

func New(options ...Option) IZLogger {
	log := new(ZLogger)
	log.opts = loadOptions(options...)
	log.init()
	return log
}
