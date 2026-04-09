package errs

import "errors"

var (
	ErrAccountAlreadyExists = errors.New("account already exists")
	ErrAccountDoesNotExists = errors.New("account does not exists")
	ErrInvalidPassword      = errors.New("invalid password")
)
