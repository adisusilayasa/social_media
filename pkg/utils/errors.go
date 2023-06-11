package utils

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
)

var (
	BadRequest          = errors.New("Bad Request")
	InternalServerError = errors.New("Internal Server Error")
	NotFound            = errors.New("Not Found")
)

type RestErr interface {
	StatusCode() int
	Error() string
	Cause() interface{}
}

type RestError struct {
	ErrStatus int         `json:"status,omitempty"`
	ErrError  string      `json:"error,omitempty"`
	ErrCause  interface{} `json:"-"`
}

func (e RestError) Error() string {
	return fmt.Sprintf("%v", e.ErrCause)
}

func (e RestError) StatusCode() int {
	return e.ErrStatus
}

func (e RestError) Cause() interface{} {
	return e.ErrCause
}

func NewRestError(status int, err string, causes interface{}) RestErr {

	return RestError{
		ErrStatus: status,
		ErrError:  err,
		ErrCause:  causes,
	}
}

func NewInternalServerError(causes interface{}) RestErr {
	result := RestError{
		ErrStatus: http.StatusInternalServerError,
		ErrError:  InternalServerError.Error(),
		ErrCause:  causes,
	}
	return result
}

func ParseError(err error) RestErr {
	switch {

	case errors.Is(err, sql.ErrNoRows):
		return NewRestError(http.StatusNotFound, NotFound.Error(), err)

	default:
		if restErr, ok := err.(RestError); ok {
			return restErr
		}
		return NewInternalServerError(err)
	}
}
