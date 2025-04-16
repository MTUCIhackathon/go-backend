package cache

import (
	"github.com/google/uuid"

	"github.com/MTUCIhackathon/go-backend/internal/model/dto"
)

type Cache interface {
	Close() error
	Set(key uuid.UUID, test dto.Test) error
	Get(key uuid.UUID) (*dto.Test, error)
	GetAll() ([]dto.Test, error)
	GetKeys() []uuid.UUID
}
