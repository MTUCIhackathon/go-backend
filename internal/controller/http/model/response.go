package model

import (
	"github.com/google/uuid"
	"time"
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
