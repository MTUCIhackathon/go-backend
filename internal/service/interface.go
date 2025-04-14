package service

import (
	"context"
	"github.com/MTUCIhackathon/go-backend/internal/model/dto"
	"github.com/google/uuid"
)

type Interface interface {
	GetAllTests(_ context.Context, token string) ([]dto.Test, error)
	GetTestByID(_ context.Context, token string, testID uuid.UUID) (*dto.Test, error)

	//CreateResolved(req dto.CreateResolved) (*dto.Resolved, error)
	//GetResolvedByUserID(userID uuid.UUID) (*dto.Resolved, error)
	//GetManyResolved(userID uuid.UUID) ([]dto.Resolved, error)

	//GetOldResolvedByID(id uuid.UUID) (*dto.Resolved, error)
	//CreateConsumer(e echo.Context, req dto.CreateConsumer) (*dto.Token, error)
	//GetConsumerByID(id uuid.UUID) (*dto.Consumer, error)
	//UpdateConsumerPassword(req dto.UpdatePassword) (bool, error)
	//SendConsumerResult(req dto.SendConsumerResult) (bool, error)
	//CreateTokensForUser(token uuid.UUID) (*dto.Token, error)

	CreateConsumer(ctx context.Context, req dto.CreateConsumer) (*dto.Token, error)
	UpdateConsumerPassword(ctx context.Context, req dto.UpdatePassword) error
	DeleteConsumerByID(ctx context.Context, token string) error
	GetConsumerByID(ctx context.Context, token string) (*dto.Consumer, error)
	Login(ctx context.Context, req dto.Login) (*dto.Token, error)
	RefreshToken(_ context.Context, token string) (*dto.Token, error)
}

// Client

// GetSummary
// GetProfessions
// GetTest
