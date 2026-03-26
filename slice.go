package util

// SliceMap returns a new slice by applying mapper to each input value.
func SliceMap[T any, R any](values []T, mapper func(T) R) []R {
	if len(values) == 0 {
		return nil
	}

	result := make([]R, 0, len(values))
	for _, value := range values {
		result = append(result, mapper(value))
	}
	return result
}

// SliceToAny converts a typed slice into []any while preserving order.
func SliceToAny[T any](values []T) []any {
	return SliceMap(values, func(value T) any {
		return value
	})
}

// SliceCompact returns a new slice with zero-value items removed while
// preserving order.
func SliceCompact[T comparable](values []T) []T {
	if len(values) == 0 {
		return nil
	}

	var zero T
	result := make([]T, 0, len(values))
	for _, value := range values {
		if value == zero {
			continue
		}
		result = append(result, value)
	}
	return result
}

// SliceCompactUnique returns a new slice with zero-value items and duplicate
// values removed while preserving the first occurrence order.
func SliceCompactUnique[T comparable](values []T) []T {
	return SliceUnique(SliceCompact(values))
}

// SliceUnique returns a new slice with duplicate values removed while preserving
// the first occurrence order.
func SliceUnique[T comparable](values []T) []T {
	if len(values) == 0 {
		return nil
	}

	result := make([]T, 0, len(values))
	seen := make(map[T]struct{}, len(values))
	for _, value := range values {
		if _, ok := seen[value]; ok {
			continue
		}
		seen[value] = struct{}{}
		result = append(result, value)
	}
	return result
}

// SliceUniqueBy returns a new slice with duplicate keys removed while
// preserving the first occurrence order.
func SliceUniqueBy[T any, K comparable](values []T, keyFn func(T) K) []T {
	if len(values) == 0 {
		return nil
	}

	result := make([]T, 0, len(values))
	seen := make(map[K]struct{}, len(values))
	for _, value := range values {
		key := keyFn(value)
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
		result = append(result, value)
	}
	return result
}
