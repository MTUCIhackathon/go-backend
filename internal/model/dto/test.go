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
