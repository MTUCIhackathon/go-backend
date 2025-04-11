package store

import (
	"context"
	"github.com/MTUCIhackathon/go-backend/internal/model/dto"
)

type (
	Interface interface {
		Forms() FormsRepository
		Consumers() ConsumersRepository
	}

	ConsumersRepository interface {
		CreateConsumer(ctx context.Context, req *dto.Consumer) error
		GetConsumerByLogin(ctx context.Context, login string) error
	}
	FormsRepository interface {
		// methods
	}
)
