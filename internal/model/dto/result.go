package dto

import (
	"time"

	"github.com/google/uuid"
)

type Result struct {
	ID            uuid.UUID
	UserID        uuid.UUID
	ResolvedID    uuid.UUID
	ImageLocation *string
	Profession    []string
	CreatedAt     time.Time
}
