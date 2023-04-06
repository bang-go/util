package util

import (
	"github.com/bang-go/util/constraint"
	"sort"
)

// SliceContainValue reports whether v is present in s.
func SliceContainValue[E comparable](s []E, v E) bool {
	return SliceIndex(s, v) >= 0
}

// SliceIndex returns the index of the first occurrence of v in s,
// or -1 if not present.
func SliceIndex[E comparable](s []E, v E) int {
	for i, vs := range s {
		if v == vs {
			return i
		}
	}
	return -1
}

func SliceSort[T constraint.Ordered](s []T) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}
