package dto

import (
	"time"

	"github.com/google/uuid"
)

type Consumer struct {
	ID        uuid.UUID
	Email     *string
	Login     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateConsumer struct {
	Login    string
	Email    *string
	Password string
}

type UpdatePassword struct {
	UserID   uuid.UUID
	Token    string
	Password string
}

type SendConsumerResult struct {
	UserID uuid.UUID
	Email  *string
}
