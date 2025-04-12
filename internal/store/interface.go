package store

import (
	"context"
	"github.com/MTUCIhackathon/go-backend/internal/model/dto"
	"github.com/google/uuid"
)

type (
	Interface interface {
		//Forms() FormsRepository
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
	FormsRepository interface {
		// methods
	}
)
