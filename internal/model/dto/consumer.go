package dto

import (
	"time"

	"github.com/google/uuid"
)

type (
	CreateConsumer struct {
		Login    string
		Password string
	}
	Consumer struct {
		ID        uuid.UUID
		Login     string
		Password  string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)

type UpdatePassword struct {
	UserID      uuid.UUID
	Token       string
	OldPassword string
	NewPassword string
}

type SendConsumerResult struct {
	UserID uuid.UUID
	Email  *string
}
