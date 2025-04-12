package dto

import "github.com/google/uuid"

type (
	ConsumerDataInToken struct {
		ID       uuid.UUID
		IsAccess bool
	}
)
