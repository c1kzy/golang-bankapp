package core_errors

import "errors"

//http
var (
	ErrInvalidArgument = errors.New("invalid argument")
	ErrConflict        = errors.New("conflict")
	ErrNotFound        = errors.New("not found")
	ErrBadRequest      = errors.New("bad request")
)

//grpc
var (
	InvalidArgumentMessage = "invalid argument"
	ConflictMessage        = "conflict"
	NotFoundMessage        = "not found"
	BadRequestMessage      = "bad request"
)
