package dto

import (
	"time"

	"github.com/google/uuid"
)

type Result struct {
	UserID     uuid.UUID
	ResolvedID uuid.UUID
	Profession []string
	CreatedAt  time.Time
}
