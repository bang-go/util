package util

// If acts as a ternary operator: returns trueVal if ok is true, else falseVal.
func If[T any](ok bool, trueVal, falseVal T) T {
	if ok {
		return trueVal
	}
	return falseVal
}

// Block blocks the current goroutine forever.
func Block() {
	select {}
}
