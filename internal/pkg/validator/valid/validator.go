package valid

import "go.uber.org/zap"

type Validator struct {
	log *zap.Logger
}

func NewValidator(log *zap.Logger) *Validator {
	if log == nil {
		log = zap.NewNop()
	}
	log.Info("validator initialize successfully")
	return &Validator{
		log: log.Named("validator"),
	}
}
