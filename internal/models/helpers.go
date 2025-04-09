package models

import "github.com/google/uuid"

type UserDataInToken struct {
	UserId   uuid.UUID `json:"user_id"`
	IsAccess bool      `json:"is_access"`
}
