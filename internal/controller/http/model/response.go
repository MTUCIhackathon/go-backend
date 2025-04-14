package model

import (
	"time"

	"github.com/google/uuid"
)

type CreateConsumerResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type (
	GetConsumerResponse struct {
		ID        uuid.UUID `json:"id"`
		Login     string    `json:"login"`
		CreatedAt time.Time `json:"created_at"`
	}
)

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type (
	GetTestResponse struct {
		ID        uuid.UUID      `json:"id"`
		Name      string         `json:"name"`
		Questions []TestQuestion `json:"questions"`
	}
	TestQuestion struct {
		Order    int    `json:"order"`
		Question string `json:"question"`
	}
)
