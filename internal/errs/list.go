package errs

import "errors"

var (
	ErrUserNotFound                    = errors.New("user not found")
	ErrEmployeeNotFound                = errors.New("employee not found")
	ErrInvalidUserID                   = errors.New("invalid user id")
	ErrNotFound                        = errors.New("not found")
	ErrInvalidRequestBody              = errors.New("invalid request body")
	ErrInvalidFieldValue               = errors.New("invalid field value")
	ErrEmployeeNameAlreadyExist        = errors.New("employee name already exist")
	ErrIncorrectEmployeeNameOrPassword = errors.New("incorrect employee name or password")
	ErrInvalidToken                    = errors.New("invalid token")
	ErrSomethingWentWrong              = errors.New("something went wrong")
)
