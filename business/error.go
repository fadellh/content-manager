package business

import "errors"

var (
	ErrInternalServerError = errors.New("internal server error")

	ErrDatabase = errors.New("something wrong in database")

	ErrNotFound = errors.New("content was not found")

	ErrInvalidSpec = errors.New("given spec is not valid")
)
