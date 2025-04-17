package dto

import (
	"github.com/google/uuid"
)

type Test struct {
	ID          uuid.UUID
	Name        string
	Description string
	Questions   []TestQuestion
}

type TestQuestion struct {
	Order    int
	Question string
}

type ThirdTestAnswers struct {
	QA map[string]string
}

type ThirdTestQuestions struct {
	Questions string   `json:"question"`
	Answers   []string `json:"answers"`
}

type QA struct {
	UserAnswers map[string]string
}
