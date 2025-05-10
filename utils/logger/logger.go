package logger

import (
	"strings"

	"go.uber.org/zap"
)

var Logger *zap.Logger

func NewLogger(level string) {
	config := zap.NewProductionConfig()
	config.Level = getLevel(strings.ToUpper(level))
	buildConfig, err := config.Build()
	if err != nil {
		panic(err)
	}

	Logger = buildConfig
}

func getLevel(level string) zap.AtomicLevel {
	switch level {
	case "DEBUG":
		return zap.NewAtomicLevelAt(zap.DebugLevel)
	case "WARN":
		return zap.NewAtomicLevelAt(zap.WarnLevel)
	case "ERROR":
		return zap.NewAtomicLevelAt(zap.ErrorLevel)
	case "INFO", "":
		return zap.NewAtomicLevelAt(zap.InfoLevel)
	default:
		return zap.NewAtomicLevelAt(zap.InfoLevel)
	}
}
