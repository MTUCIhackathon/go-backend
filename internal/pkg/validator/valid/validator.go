package valid

import "go.uber.org/zap"

type Validator struct {
	log *zap.Logger
}

func NewValidator(log *zap.Logger) *Validator {
	return &Validator{
		log: log,
	}
}
