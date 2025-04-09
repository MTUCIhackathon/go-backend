package token

import (
	"github.com/MTUCIhackathon/server/internal/models"
	"github.com/google/uuid"
)

type Provider interface {
	CreateTokenForUser(userID uuid.UUID, isAccess bool) (string, error)
	GetDataFromToken(token string) (*models.UserDataInToken, error)
}
