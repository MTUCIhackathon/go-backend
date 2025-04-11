package dto

import (
	"time"

	"github.com/google/uuid"
)

type (
	CreateConsumer struct {
		Login    string
		Email    *string
		Password string
	}
	Consumer struct {
		ID        uuid.UUID
		Login     string
		Email     *string
		Password  string
		CreatedAt time.Time
	}
)

type UpdatePassword struct {
	UserID   uuid.UUID
	Token    string
	Password string
}

type SendConsumerResult struct {
	UserID uuid.UUID
	Email  *string
}
