package service

import (
	"context"
	"github.com/MTUCIhackathon/go-backend/internal/model/dto"
	"net/http"
)

type Interface interface {
	// TODO ?
	//GetTestByName(name string) string
	//GetManyTest()

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
	UpdateConsumerPassword(ctx context.Context, r *http.Request, req dto.UpdatePassword) error
	DeleteConsumerByID(ctx context.Context, r *http.Request) error
	GetConsumerByID(ctx context.Context, r *http.Request) (*dto.Consumer, error)
	Login(ctx context.Context, req dto.Login) (*dto.Token, error)
	RefreshToken(ctx context.Context, r *http.Request) (*dto.Token, error)
}

// Client

// GetSummary
// GetProfessions
// GetTest
