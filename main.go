package main

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrInternalError    = errors.New("internal error")
	ErrAccessDenied     = errors.New("access denied")
	ErrResourceNotFound = errors.New("resource not found")
	ErrInvalidJson      = errors.New("invalid json document")
)

type Code uint32

const (
	CodeUnknownError Code = iota
	CodeInternalError
	CodeAccessDenied
	CodeResourceNotFound
	CodeInvalidJson
)

type APIError struct {
	Code        Code
	CodeString  string
	Message     string
	Explanation string
	HttpStatus  int
}

var (
	apiErrors = map[Code]APIError{
		CodeUnknownError: {
			Code:       CodeUnknownError,
			CodeString: "unknown_error",
			Message:    "unknown error",
			HttpStatus: http.StatusInternalServerError,
		},

		CodeInternalError: {
			Code:       CodeInternalError,
			CodeString: "internal_error",
			Message:    "internal error. Please, try again later",
			HttpStatus: http.StatusInternalServerError,
		},

		CodeAccessDenied: {
			Code:       CodeAccessDenied,
			CodeString: "access_denied",
			Message:    "access denied",
			HttpStatus: http.StatusForbidden,
		},

		CodeResourceNotFound: {
			Code:       CodeResourceNotFound,
			CodeString: "not_found",
			Message:    "resource not found",
			HttpStatus: http.StatusNotFound,
		},

		CodeInvalidJson: {
			Code:       CodeInvalidJson,
			CodeString: "invalid_json",
			Message:    "invalid json document",
			HttpStatus: http.StatusUnprocessableEntity,
		},
	}
)

func New(code Code) *APIError {
	e, ok := apiErrors[code]
	if !ok {
		e = apiErrors[CodeUnknownError]
	}
	return &e
}

func NewWithMessage(code Code, message string) *APIError {
	e := New(code)
	e.Message = message

	return e
}

func (ae *APIError) Error() string {
	return ae.Message
}

func convertError(err error) {
	var aErr *APIError = nil

	switch err {
	case ErrInternalError:
		aErr = New(CodeInternalError)
	case ErrAccessDenied:
		aErr = New(CodeAccessDenied)
	case ErrResourceNotFound:
		aErr = New(CodeResourceNotFound)
	case ErrInvalidJson:
		aErr = New(CodeInvalidJson)
	default:
		aErr = New(CodeUnknownError)
	}

	fmt.Println("base error: ", aErr)
}

func main() {
	convertError(ErrAccessDenied)
	convertError(ErrInvalidJson)
}
