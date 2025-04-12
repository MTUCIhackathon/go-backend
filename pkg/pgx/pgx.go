package pgx

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"

	"github.com/MTUCIhackathon/go-backend/internal/config"
)

func New(cfg config.Postgres, log *zap.Logger, afterConnect ...AfterConnect) (*pgxpool.Pool, error) {
	if log == nil {
		log = zap.L().Named("pgx")
	}

	if len(afterConnect) == 0 {
		afterConnect = []AfterConnect{AddUUIDSupport}
	}

	poolConfig, err := pgxpool.ParseConfig(cfg.GetURI())
	if err != nil {
		log.Error("failed to parse postgres config", zap.Error(err))
		return nil, errParseConfig
	}

	poolConfig.AfterConnect = load(afterConnect...)

	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		log.Error("failed to create postgres pool", zap.Error(err))
		return nil, errCreatingPool
	}

	log.Info("successfully connected to postgres pool")
	return pool, nil
}

func load(afterConnects ...AfterConnect) func(ctx context.Context, conn *pgx.Conn) error {
	return func(ctx context.Context, conn *pgx.Conn) error {
		for _, afterConnect := range afterConnects {
			if err := afterConnect(ctx, conn); err != nil {
				return err
			}
		}
		return nil
	}
}
