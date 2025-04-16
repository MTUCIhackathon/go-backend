package ml

import (
	"github.com/MTUCIhackathon/go-backend/internal/ml/client/model"
	"github.com/MTUCIhackathon/go-backend/internal/model/dto"
)

type Interface interface {
	HandlerSendResultsForFirstTest(areas []dto.Area) ([]string, error)
	HandlerSendResultsForSecondTest(kind string) (*model.PersonalityTestMLResponse, error)
}
