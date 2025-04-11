package store

import (
	"context"

	"github.com/MTUCIhackathon/go-backend/internal/model/dto"
)

type (
	Interface interface {
		Resolved() ResolvedRepository
		Consumers() ConsumersRepository
	}

	ConsumersRepository interface {
		Create(ctx context.Context, consumer dto.Consumer) error
	}
	ResolvedRepository interface {
		// methods
	}
)
