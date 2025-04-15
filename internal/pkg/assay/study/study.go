package study

import (
	"go.uber.org/zap"

	"github.com/MTUCIhackathon/go-backend/internal/pkg/assay"
)

type Study struct {
	log   *zap.Logger
	first *First
}

func New(log *zap.Logger) *Study {
	if log == nil {
		log = zap.NewNop()
		log.Warn("Study initializing failed")
	}
	log.Named("study")

	firstTest := NewFirst(log)

	return &Study{
		log:   log,
		first: firstTest,
	}
}

func (s *Study) First() assay.First {
	return s.first
}
