// Package log provides generic logging functions, decouples from external logging library
package log

import (
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

func init() {
	zapLogger, err := zap.NewDevelopmentConfig().
		Build(zap.AddCallerSkip(1))

	if err != nil {
		panic(err)
	}
	defer zapLogger.Sync()
	logger = zapLogger.Sugar()
}

func Info(args ...any) {
	logger.Info(args...)
}

func Infow(msg string, keyAndValues ...any) {
	logger.Infow(msg, keyAndValues...)
}

func Debug(args ...any) {
	logger.Debug(args...)
}

func Debugw(msg string, keyAndValues ...any) {
	logger.Debugw(msg, keyAndValues...)
}

func Warnw(msg string, keyAndValues ...any) {
	logger.Warnw(msg, keyAndValues...)
}

func Errorw(msg string, keyAndValues ...any) {
	logger.Errorw(msg, keyAndValues...)
}
