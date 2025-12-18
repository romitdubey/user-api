package logger

import (
	"go.uber.org/zap"
	// "go.uber.org/zap/zapcore"
	// "os"
)

func New(env string) *zap.Logger {
	var cfg zap.Config

	if env == "production" {
		cfg = zap.NewProductionConfig()
	} else {
		cfg = zap.NewDevelopmentConfig()
	}

	// Console output (terminal)
	cfg.OutputPaths = []string{"stdout"}
	cfg.ErrorOutputPaths = []string{"stderr"}

	// OPTIONAL: also log to file
	cfg.OutputPaths = append(cfg.OutputPaths, "logs/app.log")

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	return logger
}
