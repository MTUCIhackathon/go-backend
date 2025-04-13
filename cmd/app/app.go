package main

import (
	"github.com/MTUCIhackathon/go-backend/internal/cache/inmemory"
	"github.com/MTUCIhackathon/go-backend/internal/config"
	"github.com/MTUCIhackathon/go-backend/internal/controller/http"
	"github.com/MTUCIhackathon/go-backend/internal/pkg/encryptor/hash"
	"github.com/MTUCIhackathon/go-backend/internal/pkg/token/jwt"
	"github.com/MTUCIhackathon/go-backend/internal/pkg/validator/valid"
	"github.com/MTUCIhackathon/go-backend/internal/service/production"
	storage "github.com/MTUCIhackathon/go-backend/internal/store/pgx"
	"github.com/MTUCIhackathon/go-backend/pkg/logger"
	"github.com/MTUCIhackathon/go-backend/pkg/pgx"
)

func main() {
	log, err := logger.New("dev")
	if err != nil {
		panic(err)
	}
	log.Info("initialized config")
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}
	c, err := inmemory.New(cfg.Cache, log, inmemory.WithLoader())
	if err != nil {
		panic(err)
	}
	prv, err := jwt.NewProvider(cfg, log)
	if err != nil {
		panic(err)
	}
	pool, err := pgx.New(*cfg.Postgres, log)
	if err != nil {
		panic(err)
	}
	store, err := storage.New(log, pool)
	if err != nil {
		panic(err)
	}
	encrypt := hash.New(log)
	val := valid.NewValidator(log)
	srv, err := production.New(log, store, prv, cfg, encrypt, val, c)
	if err != nil {
		panic(err)
	}
	server, err := http.New(cfg, log, srv)
	if err != nil {
		panic(err)
	}

	err = server.Start()
	if err != nil {
		panic(err)
	}
}
