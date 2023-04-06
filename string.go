package util

import (
	"strconv"
	"strings"
)

// StringToInt   String转化为Int
func StringToInt(val string) (res int, err error) {
	return strconv.Atoi(val)
}

// StringToFloat StringToFloat  String转化为Float
func StringToFloat(val string) (res float64, err error) {
	return strconv.ParseFloat(val, 64)
}

// StringContainValue string 是否包含目标值
func StringContainValue(str string, substr string) bool {
	return strings.Contains(str, substr)
}

// IsBlank checks whether the given string is blank.
func IsBlank(s string) bool {
	return strings.TrimSpace(s) == ""
}
