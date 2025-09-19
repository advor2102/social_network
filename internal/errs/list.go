package errs

import "errors"

var (
	ErrProductNotFound  = errors.New("user not found")
	ErrInvalidProductID = errors.New("invalid user id")
	ErrNotFound         = errors.New("not found")
)
