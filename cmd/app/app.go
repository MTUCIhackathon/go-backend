package main

import (
	"context"

	"github.com/MTUCIhackathon/go-backend/internal/config"
	"github.com/MTUCIhackathon/go-backend/internal/pkg/token/jwt"
	"github.com/MTUCIhackathon/go-backend/pkg/logger"
	"github.com/MTUCIhackathon/go-backend/pkg/s3"
)

func main() {
	log, err := logger.New("dev")

	if err != nil {
		panic(err)
	}
	log.Info("initialized logger")
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}
	_, err = jwt.NewProvider(cfg, log)
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	_, err = s3.New(ctx, cfg.AWS)
}
