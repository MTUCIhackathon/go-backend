package dto

import "github.com/google/uuid"

type UserDataInToken struct {
	UserID   uuid.UUID
	IsAccess bool
}
type Token struct {
	AccessToken  string
	RefreshToken string
}
