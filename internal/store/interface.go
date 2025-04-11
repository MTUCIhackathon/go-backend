package store

import (
	"context"
	"github.com/MTUCIhackathon/go-backend/internal/model/dto"
	"github.com/google/uuid"
)

type (
	Interface interface {
		Forms() FormsRepository
		Consumers() ConsumersRepository
	}

	ConsumersRepository interface {
		CreateConsumer(ctx context.Context, req *dto.Consumer) error
		GetConsumerByLogin(ctx context.Context, login string) error
		UpdatePasswordByID(ctx context.Context, password string, id uuid.UUID) error
	}
	FormsRepository interface {
		// methods
	}
)
