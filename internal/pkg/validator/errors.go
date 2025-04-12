package validator

import "errors"

var (
	ErrorBadEmail    = errors.New("bad email")
	ErrorBadPassword = errors.New("bad password")
	ErrorRegexp      = errors.New("failed to compile regular expression")
)
