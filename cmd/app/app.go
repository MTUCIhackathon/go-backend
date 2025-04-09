package main

import (
	"github.com/MTUCIhackathon/server/internal/config"
	"github.com/MTUCIhackathon/server/internal/pkg/token/jwt"
	logger "github.com/MTUCIhackathon/server/pkg/logger"
	"github.com/google/uuid"
	"go.uber.org/zap"
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
	prov, err := jwt.NewProvider(cfg, log)
	if err != nil {
		panic(err)
	}
	id := uuid.New()
	access, err := prov.CreateTokenForUser(id, true)
	refresh, err := prov.CreateTokenForUser(id, false)
	log.Info("initialized jwt", zap.Any("access", access))
	log.Info("initialized jwt", zap.Any("refresh", refresh))

	data, err := prov.GetDataFromToken(access)
	log.Info("initialized jwt", zap.Any("data", data))
	data, err = prov.GetDataFromToken(refresh)
	log.Info("initialized jwt", zap.Any("data", data))
}
