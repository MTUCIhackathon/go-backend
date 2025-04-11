package dto

import (
	"github.com/google/uuid"
)

type Resolved struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Version   uint32
	IsActive  bool
	Questions []Question
}

type Question struct {
	ResolvedID    uuid.UUID
	QuestionOrder uint32
	Question      string
	ImageLocation *string
	Answers       []Answer
}

type Answer struct {
	ResolvedID    uuid.UUID
	QuestionOrder uint32
	OptionOrder   uint32
	Option        string
	Mark          int8
}

type CreateResolved struct {
	UserID    uuid.UUID
	Questions []Question
}
