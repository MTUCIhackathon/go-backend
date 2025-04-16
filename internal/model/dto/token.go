package dto

import (
	"time"

	"github.com/google/uuid"
)

type ConsumerDataInToken struct {
	ID        uuid.UUID
	IsAccess  bool
	ExpiresAt time.Time
	NotBefore time.Time
	IssuedAt  time.Time
}

type Token struct {
	AccessToken  string
	RefreshToken string
}
