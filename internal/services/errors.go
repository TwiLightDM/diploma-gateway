package services

import "errors"

var (
	LifetimeIsOverError          = errors.New("lifetime is over")
	UnexpectedSigningMethodError = errors.New("unexpected signing method")
)
