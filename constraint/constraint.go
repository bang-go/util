package constraint

// Complex matches the built-in complex number types and named aliases of them.
type Complex interface {
	~complex64 | ~complex128
}

// Float matches the built-in floating-point types and named aliases of them.
type Float interface {
	~float32 | ~float64
}

// String matches the built-in string type and named aliases of it.
type String interface {
	~string
}

// Signed matches the built-in signed integer types and named aliases of them.
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned matches the built-in unsigned integer types and named aliases of them.
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Integer matches the built-in signed and unsigned integer types.
type Integer interface {
	Signed | Unsigned
}

// Ordered matches types that support the standard ordering operators.
type Ordered interface {
	Integer | Float | String
}
