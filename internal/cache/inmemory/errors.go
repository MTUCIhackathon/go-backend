package inmemory

import (
	"errors"
)

var (
	ErrReadingFile   = errors.New("failed to read file")
	ErrUnmarshalling = errors.New("error while unmarshalling into structure")
	ErrNilConfig     = errors.New("failed to initialize cache: provided nil config")
	ErrNilReference  = errors.New("failed to apply option: nil reference")
)
