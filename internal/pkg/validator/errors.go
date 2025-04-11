package validator

import "errors"

var (
	ErrorBadEmail = errors.New("bad email")
	ErrorRegexp   = errors.New("failed to compile regular expression")
)
