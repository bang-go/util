package util

import (
	"errors"
	"math"
	"strconv"

	"github.com/bang-go/util/constraint"
	"github.com/shopspring/decimal"
)

// ErrDivisionByZero indicates that a float division helper was called with a zero divisor.
var ErrDivisionByZero = errors.New("util: division by zero")

// FloatToString formats val using fixed-point notation with prc digits after the decimal point.
func FloatToString[T constraint.Float](val T, prc int) string {
	return strconv.FormatFloat(float64(val), 'f', prc, floatBitSize[T]())
}

// FloatAdd returns the sum of args.
func FloatAdd[T constraint.Float](args ...T) T {
	var result decimal.Decimal
	for _, v := range args {
		result = result.Add(decimal.NewFromFloat(float64(v)))
	}
	return decimalToFloat[T](result)
}

// FloatSub subtracts args from minuend.
func FloatSub[T constraint.Float](minuend T, args ...T) T {
	result := decimal.NewFromFloat(float64(minuend))
	for _, v := range args {
		result = result.Sub(decimal.NewFromFloat(float64(v)))
	}
	return decimalToFloat[T](result)
}

// FloatMul returns the product of arg1 and args.
func FloatMul[T constraint.Float](arg1 T, args ...T) T {
	result := decimal.NewFromFloat(float64(arg1))
	for _, v := range args {
		result = result.Mul(decimal.NewFromFloat(float64(v)))
	}
	return decimalToFloat[T](result)
}

// FloatDiv divides dividend by args and returns ErrDivisionByZero when any divisor is zero.
func FloatDiv[T constraint.Float](dividend T, args ...T) (T, error) {
	result := decimal.NewFromFloat(float64(dividend))
	for _, v := range args {
		if v == 0 {
			var zero T
			return zero, ErrDivisionByZero
		}
		result = result.Div(decimal.NewFromFloat(float64(v)))
	}
	return decimalToFloat[T](result), nil
}

// MustFloatDiv divides dividend by args and panics when any divisor is zero.
func MustFloatDiv[T constraint.Float](dividend T, args ...T) T {
	res, err := FloatDiv(dividend, args...)
	if err != nil {
		panic(err)
	}
	return res
}

// FloatCeil rounds val up at the given decimal precision.
func FloatCeil[T constraint.Float](val T, precision int32) T {
	v := decimal.NewFromFloat(float64(val))
	exp := decimal.New(1, precision)
	return decimalToFloat[T](v.Mul(exp).Ceil().Div(exp))
}

// FloatFloor rounds val down at the given decimal precision.
func FloatFloor[T constraint.Float](val T, precision int32) T {
	v := decimal.NewFromFloat(float64(val))
	exp := decimal.New(1, precision)
	return decimalToFloat[T](v.Mul(exp).Floor().Div(exp))
}

// FloatTruncate truncates val at the given decimal precision.
func FloatTruncate[T constraint.Float](val T, precision int32) T {
	v := decimal.NewFromFloat(float64(val))
	return decimalToFloat[T](v.Truncate(precision))
}

// FloatCompare compares v1 and v2.
func FloatCompare[T constraint.Float](v1 T, v2 T) int {
	cV1 := decimal.NewFromFloat(float64(v1))
	cV2 := decimal.NewFromFloat(float64(v2))
	return cV1.Cmp(cV2)
}

// FloatAbs returns the absolute value of v.
func FloatAbs[T constraint.Float](v T) T {
	return T(math.Abs(float64(v)))
}

func decimalToFloat[T constraint.Float](value decimal.Decimal) T {
	res, _ := value.Float64()
	return T(res)
}

func floatBitSize[T constraint.Float]() int {
	var zero T
	switch any(zero).(type) {
	case float32:
		return 32
	default:
		return 64
	}
}
