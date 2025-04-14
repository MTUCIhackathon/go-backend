package valid

import (
	"regexp"

	"go.uber.org/zap"

	"github.com/MTUCIhackathon/go-backend/internal/pkg/validator"
)

func (v *Validator) ValidatePassword(password string) error {
	if len(password) < 8 {
		v.log.Warn("Password too short", zap.String("password", password))
		return validator.ErrorLength
	}

	rgxp, err := regexp.Compile(`\d`)
	if err != nil {
		v.log.Warn("failed to compile regex", zap.Error(err))
		return validator.ErrorRegexp
	}

	v.log.Debug("regex compiled", zap.String("regex", rgxp.String()))

	if !rgxp.MatchString(password) {
		v.log.Warn("invalid password", zap.String("password", password))
		return validator.ErrorBadPassword
	}
	v.log.Debug("valid password", zap.String("password", password))
	return nil
}

func (v *Validator) ValidateEmail(email string) error {
	rgxp, err := regexp.Compile(`[a-zA-Z0-9._-]+@[a-zA-Z0-9._-]+\.[a-zA-Z0-9_-]+`)
	if err != nil {
		v.log.Debug("failed to compile regex", zap.Error(err))
		return validator.ErrorRegexp
	}

	v.log.Debug("regex compiled", zap.String("regex", rgxp.String()))

	if !rgxp.MatchString(email) {
		v.log.Debug("invalid email", zap.String("email", email))
		return validator.ErrorBadEmail
	}
	v.log.Debug("valid email", zap.String("email", email))
	return nil
}
