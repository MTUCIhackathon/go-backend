package valid

import (
	"github.com/MTUCIhackathon/go-backend/internal/pkg/validator"
	"regexp"
)

func (v *Validator) ValidatePassword(password string) error {
	rgxp, err := regexp.Compile(`^(?=.*\d)(?=.*[a-zA-Z]).{8,}$`)
	if err != nil {
		return validator.ErrorRegexp
	}

	if !rgxp.MatchString(password) {
		return validator.ErrorBadPassword
	}

	return nil
}

func (v *Validator) ValidateEmail(email string) error {
	rgxp, err := regexp.Compile(`[a-zA-Z0-9._-]+@[a-zA-Z0-9._-]+\.[a-zA-Z0-9_-]+`)
	if err != nil {
		return validator.ErrorRegexp
	}
	if !rgxp.MatchString(email) {
		return validator.ErrorBadEmail
	}

	return nil
}
