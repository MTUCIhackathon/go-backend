package pgx

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

func New(cfg Postgres, log *zap.Logger, afterConnect ...AfterConnect) (pool *pgxpool.Pool, err error) {
	if log == nil {
		log = zap.L().Named("pgx")
	}

}
