package errs

import "errors"

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidUserID      = errors.New("invalid user id")
	ErrNotFound           = errors.New("not found")
	ErrInvalidRequestBody = errors.New("invalid request body")
)
