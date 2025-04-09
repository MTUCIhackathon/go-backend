package token

import (
	"github.com/MTUCIhackathon/server/internal/models"
	"github.com/google/uuid"
)

type Provider interface {
	CreateAccessTokenForUser(userID uuid.UUID) (string, error)
	CreateRefreshTokenForUser(userID uuid.UUID) (string, error)
	GetDataFromToken(token string) (*models.UserDataInToken, error)
}
