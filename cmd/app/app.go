package main

import (
	"github.com/MTUCIhackathon/go-backend/internal/config"
	"github.com/MTUCIhackathon/go-backend/internal/controller/smtp/client"
	"github.com/MTUCIhackathon/go-backend/internal/pkg/token/jwt"
	"github.com/MTUCIhackathon/go-backend/pkg/logger"
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

	_, err = client.New(cfg, log)
	if err != nil {
		panic(err)
	}

}
