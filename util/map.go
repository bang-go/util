package util

// MapIndex returns the index of the first occurrence of v in s,
// or "" if not present.
func MapIndex[E comparable](s map[string]E, v E) string {
	for key, vs := range s {
		if v == vs {
			return key
		}
	}
	return ""
}

// MapContainKey map是否包含目标key
func MapContainKey[E comparable](s map[string]E, key string) bool {
	if _, ok := s[key]; ok {
		return true
	}
	return false
}

// MapContainValue map是否包含目标value
func MapContainValue[E comparable](s map[string]E, v E) bool {
	return len(MapIndex(s, v)) > 0
}
