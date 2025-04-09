package token

import (
	"github.com/MTUCIhackathon/go-backend/internal/model/DTO"
	"github.com/google/uuid"
)

type Provider interface {
	CreateAccessTokenForUser(userID uuid.UUID) (string, error)
	CreateRefreshTokenForUser(userID uuid.UUID) (string, error)
	GetDataFromToken(jwtToken string) (*DTO.UserDataInToken, error)
}
