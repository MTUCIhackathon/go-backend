package client

import (
	"go.uber.org/zap"

	"github.com/MTUCIhackathon/go-backend/internal/config"
)

type SMTP struct {
	log *zap.Logger
	cfg *config.Config
}

func New(cfg *config.Config, log *zap.Logger) (*SMTP, error) {
	client := &SMTP{
		log: log.Named("smtp"),
		cfg: cfg,
	}

	client.log.Info("creating SMTP client with config", zap.Any("config", cfg.SMTP))

	return client, nil
}
