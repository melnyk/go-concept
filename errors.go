package concept

import (
	"errors"
)

var (
	// ErrNotImplemented is returned when a method is not implemented
	ErrNotImplemented = errors.New("not implemented")
	// ErrKeyNotFound is returned when a key is not found
	ErrKeyNotFound = errors.New("key not found")
)
