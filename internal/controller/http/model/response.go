package model

import (
	"time"

	"github.com/google/uuid"

	"github.com/MTUCIhackathon/go-backend/internal/pkg/style/kind"
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

type (
	CreateResolvedResponse struct {
		ID        uuid.UUID                          `json:"id"`
		Questions []QuestionInCreateResolvedResponse `json:"questions"`
	}

	QuestionInCreateResolvedResponse struct {
		QuestionOrder  uint32 `json:"question_order"`
		Question       string `json:"question"`
		QuestionAnswer string `json:"question_answer"`
		Mark           int8   `json:"mark"`
	}

	GetResolvedResponse struct {
		ID           uuid.UUID          `json:"id"`
		UserID       uuid.UUID          `json:"user_id"`
		ResolvedType kind.Type          `json:"resolved_type"`
		IsActive     bool               `json:"is_active"`
		PassedAt     time.Time          `json:"passed_at"`
		Questions    []QuestionResponse `json:"questions"`
	}
	QuestionResponse struct {
		ResolvedID     uuid.UUID `json:"resolved_id"`
		QuestionOrder  uint32    `json:"question_order"`
		Issue          string    `json:"issue"`
		QuestionAnswer string    `json:"question_answer"`
		ImageLocation  *string   `json:"image_location"`
		Mark           int8      `json:"mark"`
	}
)

type (
	CreateResultResponse struct {
		ID            uuid.UUID `json:"id"`
		ResolvedID    uuid.UUID `json:"resolved_id"`
		ImageLocation *string   `json:"image_location"`
		Professions   []string  `json:"professions"`
	}
)

type (
	CreateRequestForThirstTestResponse struct {
		Questions string   `json:"questions"`
		Answers   []string `json:"answers"`
	}
)
