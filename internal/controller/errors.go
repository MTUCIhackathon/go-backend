package controller

import "errors"

var (
	ErrorReadPassword = errors.New("failed to read password error")
	ErrorReadLogin    = errors.New("failed to read login error")
)

var (
	ErrUnknown      = errors.New("")
	ErrBadRequest   = errors.New("")
	ErrNotFound     = errors.New("")
	ErrAlreadyExist = errors.New("")
	ErrUnauthorized = errors.New("")
	ErrForbidden    = errors.New("")
	ErrInternal     = errors.New("")
)
