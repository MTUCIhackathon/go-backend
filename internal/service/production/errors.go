package production

import (
	"errors"
)

var (
	ErrNilReference      = errors.New("nil reference")
	ErrEncryptedPassword = errors.New("failed to encrypt password")
)
