package assay

import (
	"github.com/MTUCIhackathon/go-backend/internal/model/dto"
)

type Interface interface {
	First() First
}

type First interface {
	GetAreas(marks []dto.Mark) ([]dto.Area, error)
}
