package dto

import (
	"time"

	"github.com/google/uuid"

	"github.com/MTUCIhackathon/go-backend/internal/pkg/style/kind"
)

type Question struct {
	ResolvedID     uuid.UUID
	QuestionOrder  uint32
	Issue          string
	QuestionAnswer string
	ImageLocation  *string
	Mark           int8
}

type Resolved struct {
	ID           uuid.UUID
	UserID       uuid.UUID
	ResolvedType kind.Type
	IsActive     bool
	CreatedAt    time.Time
	PassedAt     time.Time
	Questions    []Question
}

type QuestionRequest struct {
	ResolvedID     uuid.UUID
	QuestionOrder  uint32
	Issue          string
	QuestionAnswer string
	ImageLocation  *string
}

type ResolvedRequest struct {
	ID           uuid.UUID
	UserID       uuid.UUID
	ResolvedType kind.Type
	IsActive     bool
	CreatedAt    time.Time
	PassedAt     time.Time
	Questions    []QuestionRequest
}
