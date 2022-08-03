package usecases

import (
	"errors"
	"fmt"
	"net/http"
)

type ErrorType uint

const (
	UnknownError ErrorType = iota
	BadRequestError
	ForbiddenError
	NotFoundError
	ConflictError
	endOfError
)

type errorProp struct {
	t ErrorType
	status int
}

var errorProps = [endOfError]errorProp{
	{t: UnknownError, status: http.StatusInternalServerError},
	{t: BadRequestError, status: http.StatusBadRequest},
	{t: ForbiddenError, status: http.StatusForbidden},
	{t: NotFoundError, status: http.StatusNotFound},
	{t: ConflictError, status: http.StatusConflict},
}

type Error struct {
	errorProp
	err error
}

func (et ErrorType) New(format string, a ...interface{}) error {
	return et.Wrap(fmt.Errorf(format, a...))
}

func (et ErrorType) Wrap(err error) error {
	if et > endOfError {
		et = UnknownError
	}

	return &Error{errorProp: errorProps[et], err: err}
}

func (ep errorProp) Type() ErrorType {
	return ep.t
}

func (ep errorProp) Status() int {
	return ep.status
}

func (e *Error) Error() string {
	return e.err.Error()
}

func (e *Error) Unwrap() error {
	return e.err
}

func (e *Error) Is(target error) bool {
	t, ok := target.(*Error)
	if !ok {
		return false
	}

	return e.Type() == t.Type() && e.Error() == t.Error()
}

type errorWithType interface {
	Type() ErrorType
}

type errorWithStatus interface {
	Status() int
}

func GetErrorType(err error) ErrorType {
	switch e := err.(type) {
	case errorWithType:
		return e.Type()
	default:
		return UnknownError
	}
}

func GetErrorStatus(err error) int {
	for {
		e, ok := err.(errorWithStatus)
		if ok {
			return e.Status()
		}

		err = errors.Unwrap(err)
		if err == nil {
			return http.StatusInternalServerError
		}
	}
}
