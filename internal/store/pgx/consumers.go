package pgx

import (
	"context"
	"fmt"
	"github.com/MTUCIhackathon/go-backend/internal/model/dto"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type ConsumersRepository struct {
	log *zap.Logger
	pgx *pgxpool.Pool
}

func newConsumersRepository(log *zap.Logger, pgx *pgxpool.Pool) *ConsumersRepository {
	return &ConsumersRepository{
		log: log,
		pgx: pgx,
	}
}

func (r *ConsumersRepository) CreateConsumer(ctx context.Context, req *dto.Consumer) error {
	const createConsumer = `INSERT INTO consumer VALUES (id, login, email, password, created_at)VALUES ($1, $2, $3, $4, $5);`
	_, err := r.pgx.Exec(ctx, createConsumer, req.ID, req.Login, req.Password, req.Email, req.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *ConsumersRepository) GetConsumerByLogin(ctx context.Context, login string) error {
	const getConsumer = `SELECT * FROM consumer WHERE login = $1;`

	row := r.pgx.QueryRow(ctx, getConsumer, login)

	if row != nil {

		return fmt.Errorf("user with choosen login already exists")
	}
	return nil
}
