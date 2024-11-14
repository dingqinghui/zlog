/**
 * @Author: dingQingHui
 * @Description:
 * @File: logger
 * @Version: 1.0.0
 * @Date: 2024/8/29 15:28
 */

package zlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type ZLogger struct {
	opts        *Options
	logger      *zap.Logger
	sugarLogger *zap.SugaredLogger
	loglevel    zap.AtomicLevel
}

func (z *ZLogger) init() {
	z.loglevel = zap.NewAtomicLevelAt(z.opts.getLevel())
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "M",                                                            // 结构化（json）输出：msg的key
		LevelKey:       "L",                                                            // 结构化（json）输出：日志级别的key（INFO，WARN，ERROR等）
		TimeKey:        "T",                                                            // 结构化（json）输出：时间的key
		CallerKey:      "C",                                                            // 结构化（json）输出：打印日志的文件对应的Key
		NameKey:        "N",                                                            // 结构化（json）输出: 日志名
		StacktraceKey:  "S",                                                            // 结构化（json）输出: 堆栈
		LineEnding:     zapcore.DefaultLineEnding,                                      // 换行符
		EncodeLevel:    zapcore.LowercaseLevelEncoder,                                  // 将日志级别转换成大写（INFO，WARN，ERROR等）
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006/01/02 15:04:05.000000Z0700"), // 日志时间的输出样式
		EncodeDuration: zapcore.SecondsDurationEncoder,                                 // 消耗时间的输出样式
		EncodeCaller:   zapcore.ShortCallerEncoder,                                     // 采用短文件路径编码输出（test/main.go:14 ）
	}

	// 获取io.Writer的实现
	loggerWriter := z.opts.getWriter()
	// 实现多个输出
	var cores []zapcore.Core
	cores = append(cores, zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.AddSync(loggerWriter), z.opts.getLevel()))
	if z.opts.getPrintConsole() {
		// 同时将日志输出到控制台，NewJSONEncoder 是结构化输出
		cores = append(cores, zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), z.opts.getLevel()))
	}
	mulCore := zapcore.NewTee(cores...)
	// 设置初始化字段
	var options = []zap.Option{
		zap.AddCaller(),
		zap.AddStacktrace(zap.DPanicLevel),
		zap.AddCallerSkip(1),
	}
	options = append(options, z.opts.getZapOption()...)
	z.logger = zap.New(mulCore, options...)
	z.sugarLogger = z.logger.Sugar()

}

func (z *ZLogger) Stop() {

}

func (z *ZLogger) Debug(msg string, fields ...zap.Field) {
	if z.logger == nil {
		return
	}
	z.logger.Debug(msg, fields...)
}

func (z *ZLogger) Info(msg string, fields ...zap.Field) {
	if z.logger == nil {
		return
	}
	z.logger.Info(msg, fields...)
}

func (z *ZLogger) Warn(msg string, fields ...zap.Field) {
	if z.logger == nil {
		return
	}
	z.logger.Warn(msg, fields...)
}

func (z *ZLogger) Error(msg string, fields ...zap.Field) {
	if z.logger == nil {
		return
	}
	z.logger.Error(msg, fields...)
}

func (z *ZLogger) Panic(msg string, fields ...zap.Field) {
	if z.logger == nil {
		return
	}
	z.logger.DPanic(msg, fields...)
}

func (z *ZLogger) Fatal(msg string, fields ...zap.Field) {
	if z.logger == nil {
		return
	}
	z.logger.Fatal(msg, fields...)
}
func (z *ZLogger) SetLogLevel(logLevel zapcore.Level) {
	z.loglevel.SetLevel(logLevel)
}

func (z *ZLogger) GetLogLevel() zapcore.Level {
	return z.loglevel.Level()
}
