package service

import (
	"github.com/MTUCIhackathon/go-backend/internal/model/dto"
	"github.com/labstack/echo/v4"
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

	CreateConsumer(e echo.Context, req dto.CreateConsumer) (*dto.Token, error)
}

// Client

// GetSummary
// GetProfessions
// GetTest
