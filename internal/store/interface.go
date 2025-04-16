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
		Results() ResultsRepository
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
		CreateResolved(ctx context.Context, data dto.Resolved) error
		GetAllActiveResolvedByUserID(ctx context.Context, id uuid.UUID) ([]dto.Resolved, error)
		GetResolvedByUserID(ctx context.Context, id uuid.UUID, resolved_type string, isActive bool) (*dto.Resolved, error)
		GetResolvedByID(ctx context.Context, id uuid.UUID) (*dto.Resolved, error)
	}

	ResultsRepository interface {
		GetLastResultByResolvedID(ctx context.Context, userID uuid.UUID, formID uuid.UUID) (*dto.Result, error)
		GetLastResults(ctx context.Context, userID uuid.UUID) ([]dto.Result, error)
		DeleteResult(ctx context.Context, resultID uuid.UUID) error
		CreateResult(ctx context.Context, result dto.Result) error
	}
)
