package util

// Ptr returns a pointer to value.
func Ptr[T any](value T) *T {
	return &value
}

// ClonePtr clones the value referenced by src.
func ClonePtr[T any](src *T) *T {
	if src == nil {
		return nil
	}

	v := *src
	return &v
}

// MustDeref returns the value referenced by ptr and panics when ptr is nil.
func MustDeref[T any](value *T) T {
	return *value
}

// DerefOr returns the value referenced by ptr or defaultValue when ptr is nil.
func DerefOr[T any](value *T, defaultValue T) T {
	if value == nil {
		return defaultValue
	}

	return *value
}

// DerefZero returns the value referenced by ptr or the zero value when ptr is nil.
func DerefZero[T any](value *T) T {
	var zero T
	return DerefOr(value, zero)
}
