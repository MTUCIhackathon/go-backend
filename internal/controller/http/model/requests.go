package model

import (
	"github.com/google/uuid"
)

type (
	CreateConsumerRequest struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	UpdatePasswordRequest struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}
	LoginRequest struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
)

type (
	GetTestRequest struct {
		Name string `json:"name"`
	}
)

type (
	CreateResultRequest struct {
		ResolvedID    uuid.UUID `json:"resolved_id"`
		ImageLocation *string   `json:"image_location"`
		Professions   []string  `json:"professions"`
	}
)
