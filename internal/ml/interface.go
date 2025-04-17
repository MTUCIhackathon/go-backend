package ml

import (
	"github.com/MTUCIhackathon/go-backend/internal/ml/client/model"
	"github.com/MTUCIhackathon/go-backend/internal/model/dto"
)

type Interface interface {
	HandlerSendResultsForFirstTest(areas []dto.Area) ([]string, error)
	HandlerSendResultsForSecondTest(kind string) (*model.PersonalityTestMLResponse, error)
	HandlerSendResultsForThirdTest(questions dto.ThirdTestAnswers) (*dto.ThirdTestQuestions, error)
	HandlerGetResultByThirdTest(qa map[string]string) ([]string, error)
	HandlerGetCommonResultByML(professions [][]string) ([]string, error)
	HandlerGenerateImage(profession string) ([]byte, error)
}
