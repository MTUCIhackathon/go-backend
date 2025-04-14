package token

import (
	"github.com/MTUCIhackathon/go-backend/internal/model/dto"
	"github.com/google/uuid"
)

type Provider interface {
	CreateAccessTokenForUser(userID uuid.UUID) (string, error)
	CreateRefreshTokenForUser(userID uuid.UUID) (string, error)
	GetDataFromToken(raw string) (*dto.UserDataInToken, error)
}
