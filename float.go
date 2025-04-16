package util

import (
	"github.com/bang-go/util/constraint"
	"github.com/shopspring/decimal"
	"strconv"
)

// FloatToString String FloatToString Float转化String
// prc precision精度
func FloatToString[T constraint.Float](val T, prc int) (str string) {
	str = strconv.FormatFloat(float64(val), 'f', prc, 64)
	return
}

// FloatAdd 加法
func FloatAdd[T constraint.Float](args ...T) float64 {
	var result decimal.Decimal
	for _, v := range args {
		result = result.Add(decimal.NewFromFloat(float64(v)))
	}
	res, _ := result.Float64()
	return res
}

// FloatSub 减法
func FloatSub[T constraint.Float](minuend T, args ...T) float64 {
	result := decimal.NewFromFloat(float64(minuend))
	for _, v := range args {
		result = result.Sub(decimal.NewFromFloat(float64(v)))
	}
	res, _ := result.Float64()
	return res
}

// FloatMul 乘法
func FloatMul[T constraint.Float](arg1 T, args ...T) float64 {
	result := decimal.NewFromFloat(float64(arg1))
	for _, v := range args {
		result = result.Mul(decimal.NewFromFloat(float64(v)))
	}
	res, _ := result.Float64()
	return res
}

// FloatDiv 除法
func FloatDiv[T constraint.Float](dividend T, args ...T) float64 {
	result := decimal.NewFromFloat(float64(dividend))
	for _, v := range args {
		result = result.Div(decimal.NewFromFloat(float64(v)))
	}
	res, _ := result.Float64()
	return res
}

// FloatCeil 向下取整
func FloatCeil[T constraint.Float](val T, precision int32) int64 {
	v := decimal.NewFromFloat(float64(val))
	return v.Ceil().IntPart()
}
func FloatFloor[T constraint.Float](val T, precision int32) int64 {
	v := decimal.NewFromFloat(float64(val))
	return v.Floor().IntPart()
}

// FloatTruncate 截断
func FloatTruncate[T constraint.Float](val T, precision int32) float64 {
	v := decimal.NewFromFloat(float64(val))
	res, _ := v.Truncate(precision).Float64()
	return res
}

// FloatCompare 比较
func FloatCompare[T constraint.Float](v1 T, v2 T) int {
	cV1 := decimal.NewFromFloat(float64(v1))
	cV2 := decimal.NewFromFloat(float64(v2))
	return cV1.Cmp(cV2)
}

// FloatAbs 　绝对值
func FloatAbs[T constraint.Float](v T) float64 {
	result := decimal.NewFromFloat(float64(v)).Abs()
	res, _ := result.Float64()
	return res
}
