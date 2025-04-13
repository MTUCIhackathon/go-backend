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
		Email     *string
		Login     string
		Password  string
		CreatedAt time.Time
	}
)

type UpdatePassword struct {
	UserID      uuid.UUID
	OldPassword string
	NewPassword string
}

type SendConsumerResult struct {
	UserID uuid.UUID
	Email  *string
}

type Login struct {
	Login    string
	Password string
}
