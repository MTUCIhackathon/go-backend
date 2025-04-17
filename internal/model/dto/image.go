package dto

import (
	"github.com/google/uuid"
)

type ImageCreation struct {
	ResultID   uuid.UUID
	Profession string
}

type Image struct {
	ResultID      uuid.UUID
	Profession    string
	ImageLocation string
}
