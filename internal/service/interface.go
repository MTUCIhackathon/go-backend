package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/MTUCIhackathon/go-backend/internal/model/dto"
)

type Interface interface {
	GetAllTests(_ context.Context, token string) ([]dto.Test, error)
	GetTestByID(_ context.Context, token string, testID uuid.UUID) (*dto.Test, error)

	//GetResolvedByUserID(userID uuid.UUID) (*dto.Resolved, error)
	//GetManyResolved(userID uuid.UUID) ([]dto.Resolved, error)

	//GetOldResolvedByID(id uuid.UUID) (*dto.Resolved, error)

	//SendConsumerResult(req dto.SendConsumerResult) (bool, error)

	CreateConsumer(ctx context.Context, req dto.CreateConsumer) (*dto.Token, error)
	UpdateConsumerPassword(ctx context.Context, req dto.UpdatePassword) error
	DeleteConsumerByID(ctx context.Context, token string) error
	GetConsumerByID(ctx context.Context, token string) (*dto.Consumer, error)
	Login(ctx context.Context, req dto.Login) (*dto.Token, error)
	RefreshToken(_ context.Context, token string) (*dto.Token, error)

	PassTest(ctx context.Context, token string, req dto.ResolvedCreation) (*dto.Result, error)

	SaveResult(ctx context.Context, token string, req dto.ResultCreation) error
	GetResultByResolvedID(ctx context.Context, token string, resultID uuid.UUID) (*dto.Result, error)
	GetResultsByUserID(ctx context.Context, token string) ([]dto.Result, error)
}

// Client

// GetSummary
// GetProfessions
// GetTest
