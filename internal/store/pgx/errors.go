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
	ErrUniqueViolation  = errors.New("try put null in not null type")
	ErrCheckViolation   = errors.New("check violation")
	ErrInvalidText      = errors.New("invalid data type")
	ErrZeroReturnedRows = errors.New("zero rows returned")
)
