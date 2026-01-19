package services

import "errors"

var (
	ErrLifetimeIsOver          = errors.New("lifetime is over")
	ErrUnexpectedSigningMethod = errors.New("unexpected signing method")
)
