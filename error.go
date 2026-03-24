package util

import (
	"errors"
	"log"
	"runtime/debug"
	"time"
)

// ErrNilRetryFunc indicates that a retry helper was called with a nil function.
var ErrNilRetryFunc = errors.New("util: nil retry function")

// RecoverAndLogPanic recovers from a panic and logs the panic value with a stack trace.
func RecoverAndLogPanic() {
	if recovered := recover(); recovered != nil {
		log.Printf("panic recovered: %v\n%s", recovered, debug.Stack())
	}
}

// RunAndRecover runs f and recovers from any panic it raises.
func RunAndRecover(f func()) {
	if f == nil {
		return
	}

	defer RecoverAndLogPanic()
	f()
}

// RetryFunc defines the work executed by retry helpers.
type RetryFunc func() error

// Retry retries rf up to attempts times.
func Retry(attempts int, rf RetryFunc) error {
	return RetryWithInterval(attempts, 0, rf)
}

// RetryWithInterval retries rf up to attempts times and sleeps between failed attempts.
func RetryWithInterval(attempts int, interval time.Duration, rf RetryFunc) (err error) {
	if attempts <= 0 {
		return nil
	}

	if rf == nil {
		return ErrNilRetryFunc
	}

	for i := range attempts {
		err = rf()
		if err == nil {
			return nil
		}
		if interval > 0 && i < attempts-1 {
			time.Sleep(interval)
		}
	}
	return err
}

// FirstError returns the first non-nil error.
func FirstError(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}
