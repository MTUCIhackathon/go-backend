package pgx

import (
	"context"

	"go.uber.org/zap"

	"github.com/MTUCIhackathon/go-backend/internal/model/dto"
)

type ConsumersRepository struct {
	log   *zap.Logger
	store *Store
}

func newConsumersRepository(store *Store) *ConsumersRepository {
	return &ConsumersRepository{
		log:   store.log.Named("consumers-repository"),
		store: store,
	}
}

func (c *ConsumersRepository) Create(ctx context.Context, consumer dto.Consumer) error {
	const query = `INSERT INTO consumers(id, email, login, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`
	commandTag, err := c.store.pool.Exec(ctx, query,
		consumer.ID,
		consumer.Email,
		consumer.Login,
		consumer.Password,
		consumer.CreatedAt,
		consumer.UpdatedAt,
	)
	if err != nil {
		return c.store.pgErr(err)
	}

	if commandTag.RowsAffected() == 0 {
		return ErrZeroRowsAffected
	}
	return nil
}
