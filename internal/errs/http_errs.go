package errs

import "errors"

var (
	ErrInvalidPayload = errors.New("invalid payload")
	ErrInternalServer = errors.New("internal server error")
)
