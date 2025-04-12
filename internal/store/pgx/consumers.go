package pgx

import (
	"context"
	"github.com/google/uuid"
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
	const query = `INSERT INTO consumers(id, login, password, created_at) VALUES ($1, $2, $3, $4);`
	commandTag, err := c.store.pool.Exec(ctx, query,
		consumer.ID,
		consumer.Login,
		consumer.Password,
		consumer.CreatedAt,
	)
	if err != nil {
		c.log.Debug("failed to create consumer", zap.Error(err))
		return c.store.pgErr(err)
	}
	if commandTag.RowsAffected() == 0 {
		c.log.Debug("failed to insert consumer", zap.Error(err))
		return ErrZeroRowsAffected
	}

	c.log.Debug("successfully inserted consumer", zap.Any("consumer", consumer))
	return nil
}

func (c *ConsumersRepository) GetLoginAvailable(ctx context.Context, login string) (bool, error) {
	const query = `SELECT COUNT(*) FROM consumers WHERE login = $1;`
	var count int
	err := c.store.pool.QueryRow(ctx, query, login).Scan(&count)
	if err != nil {
		c.store.log.Debug("failed to query user", zap.Error(err))
		return false, c.store.pgErr(err)
	}

	if count != 0 {
		c.store.log.Debug("found user", zap.Int("count", count))
		return false, ErrAlreadyExists
	}

	c.store.log.Debug("login is available", zap.Any("user", login))

	return true, nil
}

func (c *ConsumersRepository) GetPasswordByID(ctx context.Context, id uuid.UUID) (string, error) {
	const query = `SELECT password FROM consumers WHERE id = $1;`
	var password string
	err := c.store.pool.QueryRow(ctx, query, id).Scan(&password)
	if err != nil {
		c.store.log.Debug("failed to query user", zap.Error(err))
		return "", c.store.pgErr(err)
	}

	if password == "" {
		c.store.log.Debug("user not found", zap.Any("user", id))
		return "", ErrZeroReturnedRows
	}

	c.store.log.Debug("user password found", zap.Any("password", password))

	return password, nil
}

func (c *ConsumersRepository) UpdatePasswordByID(ctx context.Context, id uuid.UUID, password string) error {
	const query = `UPDATE consumers SET password = $1 WHERE id = $2;`
	_, err := c.store.pool.Exec(ctx, query, password, id)
	if err != nil {
		c.store.log.Debug("failed to update user", zap.Any("user", id))
		return c.store.pgErr(err)
	}

	c.store.log.Debug("successfully updated user", zap.Any("user", id))

	return nil
}

func (c *ConsumersRepository) DeleteByID(ctx context.Context, id uuid.UUID) error {
	const query = `DELETE FROM consumers WHERE id = $1;`
	_, err := c.store.pool.Exec(ctx, query, id)
	if err != nil {
		c.store.log.Debug("failed to delete user", zap.Any("user", id))
		return c.store.pgErr(err)
	}

	c.store.log.Debug("successfully deleted user", zap.Any("user", id))

	return nil
}

func (c *ConsumersRepository) GetByID(ctx context.Context, id uuid.UUID) (*dto.Consumer, error) {
	const query = `SELECT * FROM consumers WHERE id = $1;`
	var consumer dto.Consumer
	err := c.store.pool.QueryRow(ctx, query, id).Scan(&consumer.ID, &consumer.Login, &consumer.Password, &consumer.CreatedAt)
	if err != nil {
		c.store.log.Debug("failed to query user", zap.Any("user", id))
		return nil, c.store.pgErr(err)
	}

	c.store.log.Debug("user password found", zap.Any("user", consumer))

	return &consumer, nil
}

func (c *ConsumersRepository) GetByLogin(ctx context.Context, login string) (*dto.Consumer, error) {
	const query = `SELECT * FROM consumers WHERE login = $1;`
	var consumer dto.Consumer
	err := c.store.pool.QueryRow(ctx, query, login).Scan(&consumer.ID, &consumer.Login, &consumer.Password, &consumer.CreatedAt)
	if err != nil {
		c.store.log.Debug("failed to query user", zap.Any("user", login))
		return nil, c.store.pgErr(err)
	}

	c.store.log.Debug("user password found", zap.Any("user", consumer))

	return &consumer, nil
}
