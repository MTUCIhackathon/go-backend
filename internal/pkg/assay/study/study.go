package study

import (
	"go.uber.org/zap"

	"github.com/MTUCIhackathon/go-backend/internal/pkg/assay"
	"github.com/MTUCIhackathon/go-backend/internal/pkg/assay/study/first"
)

type Study struct {
	log   *zap.Logger
	first *first.First
}

func New(log *zap.Logger) *Study {
	if log == nil {
		log = zap.NewNop()
		log.Warn("Study initializing failed")
	}
	log.Named("study")

	firstTest := first.New(log)

	return &Study{
		log:   log,
		first: firstTest,
	}
}

func (s *Study) First() assay.First {
	return s.first
}
