package pgx

import (
	"context"
	"github.com/MTUCIhackathon/go-backend/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Repository struct {
	cfg  *config.Config
	log  *zap.Logger
	pool *pgxpool.Pool
}

func NewRepository(log *zap.Logger, cfg *config.Config) (*Repository, error) {
	if log == nil {
		log = zap.NewNop()
	}
	c, err := pgxpool.ParseConfig(cfg.Postgres.GetDNS())
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	//TODO: add UUID support

	pool, err := pgxpool.NewWithConfig(context.Background(), c)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	repo := &Repository{
		cfg:  cfg,
		log:  log,
		pool: pool,
	}
	return repo, nil
}
