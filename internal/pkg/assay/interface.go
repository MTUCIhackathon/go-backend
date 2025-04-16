package assay

import (
	"github.com/MTUCIhackathon/go-backend/internal/model/dto"
)

type Interface interface {
	First() First
	Second() Second
}

type First interface {
	GetAreas(marks []dto.Mark) ([]dto.Area, error)
}

type Second interface {
	GetPersonality(marks []dto.Mark) (string, error)
}
