package util

import "github.com/bang-go/util/constraint"

// If 三元表达式
func If[T constraint.Comparable](expr bool, trueVale, falseValue T) T {
	if expr {
		return trueVale
	}
	return falseValue
}

func Block() {
	select {}
}
