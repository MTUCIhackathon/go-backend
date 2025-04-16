package model

import (
	"time"

	"github.com/google/uuid"
)

type (
	CreateConsumerResponse struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}
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
		ID          uuid.UUID      `json:"id"`
		Name        string         `json:"name"`
		Description string         `json:"description"`
		Questions   []TestQuestion `json:"questions"`
	}
	TestQuestion struct {
		Order    int    `json:"order"`
		Question string `json:"question"`
	}

	GetAllTestResponse struct {
		GetTestResponse []GetTestResponse `json:"tests"`
	}
)

type (
	GetResultResponse struct {
		ID            uuid.UUID `json:"id"`
		UserID        uuid.UUID `json:"user_id"`
		ResolvedID    uuid.UUID `json:"resolved_id"`
		ImageLocation *string   `json:"image_location"`
		Professions   []string  `json:"professions"`
		CreatedAt     time.Time `json:"created_at"`
	}
	GetResultsByUserIDResponse struct {
		Results []GetResultResponse `json:"results"`
	}
)
