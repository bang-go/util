package errcode

import (
	"errors"
	"fmt"
)

// Error represents a business error with a code and a message.
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Error implements the standard error interface.
// It returns a string representation of the error.
func (e *Error) Error() string {
	if e == nil {
		return "<nil>"
	}
	return fmt.Sprintf("code=%d message=%s", e.Code, e.Message)
}

// New creates a new Error instance.
func New(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

// As returns the first *Error found in err's chain.
func As(err error) (*Error, bool) {
	var e *Error
	if errors.As(err, &e) {
		if e != nil {
			return e, true
		}
	}
	return nil, false
}

// Code returns the error code and whether err contains an errcode.Error.
func Code(err error) (int, bool) {
	e, ok := As(err)
	if !ok {
		return 0, false
	}

	return e.Code, true
}

// Message returns the error message and whether err contains an errcode.Error.
func Message(err error) (string, bool) {
	e, ok := As(err)
	if !ok {
		return "", false
	}

	return e.Message, true
}

// Is implements the standard errors.Is interface.
// It considers two errors equal if they are both *Error and have the same Code.
// This allows errors.Is(err, ErrNotFound) to work even if messages differ.
func (e *Error) Is(target error) bool {
	if e == nil {
		return false
	}
	var t *Error
	ok := errors.As(target, &t)
	if !ok || t == nil {
		return false
	}
	return e.Code == t.Code
}

// HasCode checks if the error (or any error in its chain) is an *Error with the given code.
func HasCode(err error, code int) bool {
	e, ok := As(err)
	return ok && e.Code == code
}
