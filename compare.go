package util

import "github.com/bang-go/util/constraint"

// Min returns the smaller of a and b.
func Min[T constraint.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Max returns the larger of a and b.
func Max[T constraint.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
