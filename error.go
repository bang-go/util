package util

import (
	"log"
	"time"
)

func RecoverForPanic() {
	if err := recover(); err != nil {
		log.Printf("Unexpected panic in RecoverForPanic: %v", err)
	}
}

func RunWithRecover(f func()) {
	defer RecoverForPanic()
	f()
}

type RetryFunc = func() error

// ErrorRetry 错误重试
func ErrorRetry(maxRetry int, rf RetryFunc) (err error) {
	return ErrorRetryWithInterval(maxRetry, 0, rf)
}

func ErrorRetryWithInterval(maxRetry int, interval time.Duration, rf RetryFunc) (err error) {
	if maxRetry <= 0 {
		return
	}
	for range maxRetry {
		err = rf()
		if err == nil {
			return
		}
		time.Sleep(interval)
	}
	return
}
