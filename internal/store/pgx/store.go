package pgx

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"

	"github.com/MTUCIhackathon/go-backend/internal/store"
)

var _ store.Interface = (*Store)(nil)

type Store struct {
	log       *zap.Logger
	pool      *pgxpool.Pool
	consumers *ConsumersRepository
	resolved  *ResolvedRepository
}

func (s *Store) Resolved() store.ResolvedRepository {
	if s == nil {
		zap.L().Named("store").Named("resolved").Error(
			"got unexpectedly nil store repository",
		)
		return nil
	}

	return s.resolved
}

func (s *Store) Consumers() store.ConsumersRepository {
	if s == nil {
		zap.L().Named("store").Named("consumers").Error(
			"got unexpectedly nil store repository",
		)
		return nil
	}

	return s.consumers
}

func New(log *zap.Logger, pool *pgxpool.Pool) (*Store, error) {
	if log == nil {
		log = zap.NewNop()
	}

	if pool == nil {
		return nil, store.ErrNilPool
	}

	s := &Store{
		log:  log.Named("store"),
		pool: pool,
	}

	s.consumers = newConsumersRepository(s)
	s.resolved = newResolvedRepository(s)

	s.log.Info("store initialized successfully")
	return nil, nil
}
