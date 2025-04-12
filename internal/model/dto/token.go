package dto

import "github.com/google/uuid"

type UserDataInToken struct {
	ID       uuid.UUID `json:"id"`
	IsAccess bool      `json:"is_access"`
}
type Token struct {
	AccessToken  string
	RefreshToken string
}
