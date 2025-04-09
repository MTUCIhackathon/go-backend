package main

import (
	"github.com/MTUCIhackathon/server/internal/config"
	"github.com/MTUCIhackathon/server/internal/pkg/token/jwt"
	logger "github.com/MTUCIhackathon/server/pkg/logger"
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

}
