package pgx

import (
	"errors"
)

var (
	errCreatingPool = errors.New("failed to create postgres pool")
	errParseConfig  = errors.New("failed to parse postgres config")
)
