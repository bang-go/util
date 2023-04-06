package util

import "log"

func RecoverForPanic() {
	if err := recover(); err != nil {
		log.Printf("Unexpected panic in RecoverForPanic: %v", err)
	}
}

func RunWithRecover(f func()) {
	defer RecoverForPanic()
	f()
}
