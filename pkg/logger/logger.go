package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/MTUCIhackathon/go-backend/internal/config"
)

const (
	envDev  = "dev"
	envProd = "prod"
)

func New(conf *config.Config) (*zap.Logger, error) {
	var level zapcore.Level

	switch conf.Controller.LogLevel {
	case envDev:
		level = zap.DebugLevel
	case envProd:
		level = zap.InfoLevel
	}

	cfg := zap.Config{
		Level:       zap.NewAtomicLevelAt(level),
		Encoding:    "json",
		OutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",
			LevelKey:    "level",
			TimeKey:     "time",
			EncodeTime:  zapcore.ISO8601TimeEncoder,
			EncodeLevel: zapcore.CapitalLevelEncoder,
		},
		ErrorOutputPaths: []string{"stderr"},
	}

	log, err := cfg.Build()

	if err != nil {
		return nil, err
	}
	defer log.Sync()
	return log, nil
}
