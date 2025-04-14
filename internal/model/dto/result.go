package dto

import (
	"github.com/google/uuid"
	"time"
)

type Result struct {
	UserID     uuid.UUID
	ResolvedID uuid.UUID
	Profession []string
	CreatedAt  time.Time
}
