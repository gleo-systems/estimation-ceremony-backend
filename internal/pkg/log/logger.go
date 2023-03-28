// Application generic logger facade
package log

import (
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

func init() {
	zapLogger, _ := zap.NewDevelopment()
	defer zapLogger.Sync()
	logger = zapLogger.Sugar()
}

func Info(args ...any) {
	logger.Info(args...)
}

func Infow(msg string, keyAndValues ...any) {
	logger.Infow(msg, keyAndValues...)
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
