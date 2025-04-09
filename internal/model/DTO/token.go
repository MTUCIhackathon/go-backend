package DTO

import "github.com/google/uuid"

type UserDataInToken struct {
	UserID   uuid.UUID `json:"user_id"`
	IsAccess bool      `json:"is_access"`
}
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
