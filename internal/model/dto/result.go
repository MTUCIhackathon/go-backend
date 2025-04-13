package dto

import (
	"github.com/google/uuid"
	"time"
)

type Result struct {
	UserID      uuid.UUID
	FormID      uuid.UUID
	FormVersion uuid.UUID
	Profession  []string
	CreatedAt   time.Time
}
