package model

import "errors"

var (
	ErrInvalidRecord   = errors.New("invalid record")
	ErrOutOfStock      = errors.New("out of stock")
	ErrAlreadyArchived = errors.New("already archived")
	ErrNotFound        = errors.New("not found")
	ErrUnauthorized    = errors.New("unauthorized")
	ErrCancelled       = errors.New("cancelled")
)
