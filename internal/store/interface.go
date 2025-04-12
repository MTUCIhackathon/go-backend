package store

import (
	"context"

	"github.com/google/uuid"

	"github.com/MTUCIhackathon/go-backend/internal/model/dto"
)

type (
	Interface interface {
		Resolved() ResolvedRepository
		Consumers() ConsumersRepository
	}

	ConsumersRepository interface {
		Create(ctx context.Context, consumer dto.Consumer) error
		GetLoginAvailable(ctx context.Context, login string) (bool, error)
		GetPasswordByID(ctx context.Context, id uuid.UUID) (string, error)
		UpdatePasswordByID(ctx context.Context, id uuid.UUID, password string) error
		DeleteByID(ctx context.Context, id uuid.UUID) error
		GetByID(ctx context.Context, id uuid.UUID) (*dto.Consumer, error)
		GetByLogin(ctx context.Context, login string) (*dto.Consumer, error)
	}
	ResolvedRepository interface {
	}
)
