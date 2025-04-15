package study

import (
	"go.uber.org/zap"

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

	first := first.New(log)

	return &Study{
		log:   log,
		first: first,
	}
}

func (s *Study) First() *first.First {
	return s.first
}
