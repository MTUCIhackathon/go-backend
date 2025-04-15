package token

import (
	"github.com/google/uuid"

	"github.com/MTUCIhackathon/go-backend/internal/model/dto"
)

type Provider interface {
	CreateAccessTokenForUser(userID uuid.UUID) (string, error)
	CreateRefreshTokenForUser(userID uuid.UUID) (string, error)
	GetDataFromToken(raw string) (*dto.ConsumerDataInToken, error)
	CreateAccessAndRefreshTokenForUser(userID uuid.UUID) (string, string, error)
}
