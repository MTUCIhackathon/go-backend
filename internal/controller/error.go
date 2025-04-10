package controller

import "errors"

var (
	ErrorReadPassword = errors.New("failed to read password error")
	ErrorReadLogin    = errors.New("failed to read login error")
)
