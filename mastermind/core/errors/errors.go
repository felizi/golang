package errors

import (
	"errors"
)

var (
	ErrValidation = errors.New("validation error")
	ErrAlreadyComplete = errors.New("bufio: negative count")
)
