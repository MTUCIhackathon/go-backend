package pgx

import (
	"errors"
)

var (
	ErrAlreadyExists    = errors.New("object already exists")
	ErrNotFound         = errors.New("no founded content")
	ErrToManyContent    = errors.New("too many content to return")
	ErrUnknown          = errors.New("unknown error")
	ErrZeroRowsAffected = errors.New("zero rows affected")
)
