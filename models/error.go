package models

import (
	"errors"
	"net/http"
)

type errorCustom error

var (
	ErrorDefault      errorCustom = errors.New("Something when wrong")
	ErrorDataNotFound errorCustom = errors.New("Data not found")
)

// Error is customer handle error
type Error struct {
	Code  int
	Error error
}

func ErrorMapping(err error) Error {

	switch err {
	case ErrorDataNotFound:
		return Error{
			Code:  http.StatusNotFound,
			Error: err,
		}
	default:
		return Error{
			Code:  http.StatusInternalServerError,
			Error: ErrorDefault,
		}
	}
}
