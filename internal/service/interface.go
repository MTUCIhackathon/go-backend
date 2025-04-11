package service

import (
	"github.com/google/uuid"

	"github.com/MTUCIhackathon/go-backend/internal/model/dto"
)

type Interface interface {
	// TODO ?
	GetTestByName(name string) string
	GetManyTest()

	CreateResolved(req dto.CreateResolved) (*dto.Resolved, error)
	GetResolvedByUserID(userID uuid.UUID) (*dto.Resolved, error)
	GetManyResolved(userID uuid.UUID) ([]dto.Resolved, error)

	GetOldResolvedByID(id uuid.UUID) (*dto.Resolved, error)

	CreateConsumer(req dto.CreateConsumer) (*dto.Consumer, error)
	GetConsumerByID(id uuid.UUID) (*dto.Consumer, error)
	UpdateConsumerPassword(req dto.UpdatePassword) (bool, error)
	SendConsumerResult(req dto.SendConsumerResult) (bool, error)
}

// Client

// GetSummary
// GetProfessions
// GetTest
