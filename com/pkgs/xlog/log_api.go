package xlog

import (
	"log"

	"go.uber.org/zap"
)

func xLogger() *zap.SugaredLogger {
	if xLog == nil {
		// Shared(DefaultConfig, DefaultDirectory)
		log.Fatalln("log system not init")
	}
	return xLog
}

func Debug(args ...interface{}) {
	xLogger().Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	xLogger().Debugf(format, args...)
}

func Info(args ...interface{}) {
	xLogger().Info(args...)
}

func Infof(format string, args ...interface{}) {
	xLogger().Infof(format, args...)
}

func Warn(args ...interface{}) {
	xLogger().Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	xLogger().Warnf(format, args...)
}

func Error(args ...interface{}) {
	xLogger().Error(args...)
}

func Errorf(format string, args ...interface{}) {
	xLogger().Errorf(format, args)
}

func Panic(args ...interface{}) {
	xLogger().Panic(args...)
}

func Panicf(format string, args ...interface{}) {
	xLogger().Panicf(format, args...)
}
