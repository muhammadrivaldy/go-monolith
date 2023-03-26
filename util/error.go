package util

import (
	"errors"
	"net/http"
)

var (
	ErrorDefault      = errors.New("something when wrong")
	ErrorDataNotFound = errors.New("data not found")
	ErrorUnauthorized = errors.New("unauthorized access")
)

type Error struct {
	Code  int
	Error error
}

func (e Error) IsError() bool {
	return e.Error != nil
}

func ErrorMapping(err error) Error {

	switch err {
	case nil:
		return Error{}
	case ErrorDataNotFound:
		return Error{Code: http.StatusNotFound, Error: err}
	case ErrorUnauthorized:
		return Error{Code: http.StatusUnauthorized, Error: err}
	default:
		return Error{Code: http.StatusInternalServerError, Error: ErrorDefault}
	}
}
