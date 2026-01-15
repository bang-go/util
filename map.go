package util

// MapKeyForValue returns the first key mapped to v in m, and true if found.
// If not found, returns zero value of K and false.
// Note: Random order in maps means this might return different keys for duplicate values.
func MapKeyForValue[K comparable, V comparable](m map[K]V, v V) (K, bool) {
	var zero K
	for key, val := range m {
		if val == v {
			return key, true
		}
	}
	return zero, false
}

// MapContainValue map是否包含目标value
func MapContainValue[K comparable, V comparable](m map[K]V, v V) bool {
	for _, val := range m {
		if val == v {
			return true
		}
	}
	return false
}
