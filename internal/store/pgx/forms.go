package pgx

import (
	"go.uber.org/zap"
)

type ResolvedRepository struct {
	log   *zap.Logger
	store *Store
}

func newResolvedRepository(store *Store) *ResolvedRepository {
	return &ResolvedRepository{
		log:   store.log.Named("forms-repository"),
		store: store,
	}
}
