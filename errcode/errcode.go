package errcode

import (
	"errors"
	"fmt"
)

// Error represents a business error with a code and a message.
type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// Error implements the standard error interface.
// It returns a string representation of the error.
func (e *Error) Error() string {
	return fmt.Sprintf("code: %d, msg: %s", e.Code, e.Msg)
}

// New creates a new Error instance.
func New(code int, msg string) *Error {
	return &Error{
		Code: code,
		Msg:  msg,
	}
}

// Code returns the error code.
// This helper is useful when working with the interface error.
func Code(err error) int {
	if err == nil {
		return 0 // Or 200, depending on convention, but 0 usually means success/no error
	}
	var e *Error
	if errors.As(err, &e) {
		if e == nil {
			return -1
		}
		return e.Code
	}
	return -1 // Indicates unknown/internal error
}

// Msg returns the error message.
func Msg(err error) string {
	if err == nil {
		return ""
	}
	var e *Error
	if errors.As(err, &e) {
		if e == nil {
			return ""
		}
		return e.Msg
	}
	return err.Error()
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
	if !ok {
		return false
	}
	return e.Code == t.Code
}

// IsCode checks if the error (or any error in its chain) is an *Error with the given code.
func IsCode(err error, code int) bool {
	if err == nil {
		return false
	}
	var e *Error
	// errors.As unwraps the error chain to find the first *Error
	if errors.As(err, &e) {
		return e.Code == code
	}
	return false
}
