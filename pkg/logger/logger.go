package logger

import "go.uber.org/zap"

func CreateLogger() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	return logger.Sugar()
}
