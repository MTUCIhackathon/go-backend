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
		Level:    zap.NewAtomicLevelAt(level),
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.EpochTimeEncoder,
			EncodeDuration: zapcore.MillisDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	log, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	defer func() {
		err = log.Sync()
	}()

	zap.ReplaceGlobals(log)
	zap.RedirectStdLog(log)

	return log, nil
}
