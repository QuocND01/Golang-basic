package logger

import (
	"go.uber.org/zap"
)

var (
	Logger *zap.Logger
)

func InitLogger() {
	var err error
	Logger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
}

func Sync() {
	if Logger != nil {
		Logger.Sync()
	}
}
