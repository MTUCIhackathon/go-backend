package valid

import (
	"github.com/MTUCIhackathon/go-backend/internal/pkg/validator"
	"regexp"
)

func (v *Validator) ValidateLogin() {
}

func (v *Validator) ValidateEmail(email string) error {
	rgxp, err := regexp.Compile(`[a-zA-Z0-9._-]+@[a-zA-Z0-9._-]+\.[a-zA-Z0-9_-]+`)
	if err != nil {
		return validator.ErrorRegexp
	}
	if !rgxp.MatchString(email) {
		return validator.ErrorRegexp
	}

	return nil
}
