package store

import (
	"errors"
)

var (
	ErrNilPool  = errors.New("got unexpected nil pool")
	ErrNilStore = errors.New("nil store repository")
)
