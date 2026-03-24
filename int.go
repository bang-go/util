package util

import (
	"crypto/rand"
	"errors"
	"math/big"
	"strconv"

	"github.com/bang-go/util/constraint"
)

// ErrIntegerOutOfRange indicates that an integer result cannot be represented in the target type.
var ErrIntegerOutOfRange = errors.New("util: integer value out of range")

// IntRandRange returns a cryptographically secure random integer in the inclusive range [n1, n2].
func IntRandRange[T constraint.Integer](n1 T, n2 T) (T, error) {
	low := integerToBigInt(n1)
	high := integerToBigInt(n2)
	if high.Cmp(low) < 0 {
		low, high = high, low
	}

	span := new(big.Int).Sub(high, low)
	span.Add(span, big.NewInt(1))

	n, err := rand.Int(rand.Reader, span)
	if err != nil {
		var zero T
		return zero, err
	}

	result := new(big.Int).Add(low, n)
	return integerFromBigInt[T](result)
}

func integerToBigInt[T constraint.Integer](v T) *big.Int {
	switch value := any(v).(type) {
	case int:
		return big.NewInt(int64(value))
	case int8:
		return big.NewInt(int64(value))
	case int16:
		return big.NewInt(int64(value))
	case int32:
		return big.NewInt(int64(value))
	case int64:
		return big.NewInt(value)
	case uint:
		return new(big.Int).SetUint64(uint64(value))
	case uint8:
		return new(big.Int).SetUint64(uint64(value))
	case uint16:
		return new(big.Int).SetUint64(uint64(value))
	case uint32:
		return new(big.Int).SetUint64(uint64(value))
	case uint64:
		return new(big.Int).SetUint64(value)
	case uintptr:
		return new(big.Int).SetUint64(uint64(value))
	default:
		panic("unreachable integer type")
	}
}

func integerFromBigInt[T constraint.Integer](v *big.Int) (T, error) {
	var zero T

	switch any(zero).(type) {
	case int:
		n, ok := bigIntToInt64(v)
		if !ok {
			return zero, ErrIntegerOutOfRange
		}
		if strconv.IntSize == 32 && (n < -1<<31 || n > 1<<31-1) {
			return zero, ErrIntegerOutOfRange
		}
		return T(n), nil
	case int8:
		n, ok := bigIntToInt64(v)
		if !ok || n < -1<<7 || n > 1<<7-1 {
			return zero, ErrIntegerOutOfRange
		}
		return T(n), nil
	case int16:
		n, ok := bigIntToInt64(v)
		if !ok || n < -1<<15 || n > 1<<15-1 {
			return zero, ErrIntegerOutOfRange
		}
		return T(n), nil
	case int32:
		n, ok := bigIntToInt64(v)
		if !ok || n < -1<<31 || n > 1<<31-1 {
			return zero, ErrIntegerOutOfRange
		}
		return T(n), nil
	case int64:
		n, ok := bigIntToInt64(v)
		if !ok {
			return zero, ErrIntegerOutOfRange
		}
		return T(n), nil
	case uint:
		n, ok := bigIntToUint64(v)
		if !ok {
			return zero, ErrIntegerOutOfRange
		}
		if strconv.IntSize == 32 && n > 1<<32-1 {
			return zero, ErrIntegerOutOfRange
		}
		return T(n), nil
	case uint8:
		n, ok := bigIntToUint64(v)
		if !ok || n > 1<<8-1 {
			return zero, ErrIntegerOutOfRange
		}
		return T(n), nil
	case uint16:
		n, ok := bigIntToUint64(v)
		if !ok || n > 1<<16-1 {
			return zero, ErrIntegerOutOfRange
		}
		return T(n), nil
	case uint32:
		n, ok := bigIntToUint64(v)
		if !ok || n > 1<<32-1 {
			return zero, ErrIntegerOutOfRange
		}
		return T(n), nil
	case uint64:
		n, ok := bigIntToUint64(v)
		if !ok {
			return zero, ErrIntegerOutOfRange
		}
		return T(n), nil
	case uintptr:
		n, ok := bigIntToUint64(v)
		if !ok {
			return zero, ErrIntegerOutOfRange
		}
		if strconv.IntSize == 32 && n > 1<<32-1 {
			return zero, ErrIntegerOutOfRange
		}
		return T(n), nil
	default:
		panic("unreachable integer type")
	}
}

func bigIntToInt64(v *big.Int) (int64, bool) {
	if !v.IsInt64() {
		return 0, false
	}
	return v.Int64(), true
}

func bigIntToUint64(v *big.Int) (uint64, bool) {
	if v.Sign() < 0 || !v.IsUint64() {
		return 0, false
	}
	return v.Uint64(), true
}
