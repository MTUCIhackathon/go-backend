package model

import (
	"github.com/google/uuid"

	"github.com/MTUCIhackathon/go-backend/internal/pkg/style/kind"
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
		ResolvedID    uuid.UUID                       `json:"resolved_id"`
		TestType      kind.Type                       `json:"test_type"`
		ImageLocation *string                         `json:"image_location"`
		Questions     []QuestionInCreateResultRequest `json:"questions"`
	}
	QuestionInCreateResultRequest struct {
		QuestionOrder uint32 `json:"question_order"`
		Mark          int8   `json:"mark"`
	}
)

type (
	CreateResolvedRequest struct {
		TestType  kind.Type                         `json:"test_type"`
		Questions []QuestionInCreateResolvedRequest `json:"questions"`
	}

	QuestionInCreateResolvedRequest struct {
		QuestionOrder  uint32 `json:"question_order"`
		Question       string `json:"question"`
		QuestionAnswer string `json:"question_answer"`
	}
)
