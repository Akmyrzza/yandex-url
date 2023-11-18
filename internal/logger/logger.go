package logger

import "go.uber.org/zap"

type Logger struct {
	LoggerZap *zap.Logger
}

func NewLogger() *Logger {
	Logger := new(Logger)
	LoggerZap, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	Logger.LoggerZap = LoggerZap
	defer Logger.LoggerZap.Sync()

	return Logger
}
