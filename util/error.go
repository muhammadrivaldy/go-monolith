package util

import (
	"errors"
	"net/http"
)

var (
	ErrorDefault              = errors.New("something when wrong")
	ErrorDataNotFound         = errors.New("data not found")
	ErrorDataConflict         = errors.New("data conflict")
	ErrorUserAlreadyExists    = errors.New("user already exists")
	ErrorWrongEmailOrPassword = errors.New("wrong email or password")
	ErrorIncorrectInput       = errors.New("please put correct value")
	ErrorUnauthorized         = errors.New("unauthorized access")
	ErrorUnprocessableEntity  = errors.New("unprocessable entity")
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
	case ErrorDataConflict:
		return Error{Code: http.StatusConflict, Error: err}
	case ErrorUserAlreadyExists:
		return Error{Code: http.StatusAlreadyReported, Error: err}
	case ErrorWrongEmailOrPassword:
		return Error{Code: http.StatusUnauthorized, Error: err}
	case ErrorIncorrectInput:
		return Error{Code: http.StatusBadRequest, Error: err}
	case ErrorUnauthorized:
		return Error{Code: http.StatusUnauthorized, Error: err}
	case ErrorUnprocessableEntity:
		return Error{Code: http.StatusUnprocessableEntity, Error: err}
	default:
		return Error{Code: http.StatusInternalServerError, Error: ErrorDefault}
	}
}
