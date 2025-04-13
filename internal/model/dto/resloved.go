package dto

import (
	"github.com/google/uuid"
	"time"
)

type Question struct {
	FromID        uuid.UUID
	QuestionOrder uint32
	Issue         string
	ImageLocation *string
	Mark          int8
}

type Resolved struct {
	ID           uuid.UUID
	Version      uint32
	UserID       uuid.UUID
	ResolvedType string
	IsActive     bool
	CreatedAt    time.Time
	PassedAt     time.Time
	Questions    []Question
}
