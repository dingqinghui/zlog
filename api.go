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

func New(options ...Option) *ZLogger {
	log := new(ZLogger)
	log.opts = loadOptions(options...)
	log.init()
	return log
}

var gz IZLogger = New(WithPrintConsole(true))

func SetLogger(z IZLogger) {
	gz = z
}

func Debug(msg string, fields ...zap.Field) {
	gz.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	gz.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	gz.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	gz.Error(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	gz.Panic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	gz.Fatal(msg, fields...)
}
func SetLevel(logLevel zapcore.Level) {
	gz.SetLogLevel(logLevel)
}

func GetLevel() zapcore.Level {
	return gz.GetLogLevel()
}

func Stop() {
	gz.Stop()
}
