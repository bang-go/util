package util

import (
	"math/rand"
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

var (
	lowerLetter = "abcdefghijklmnopqrstuvwxyz"
	upperLetter = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	number      = "0123456789"
)

// StringRand 随机字符串
func StringRand(length int) string {
	charset := []byte(lowerLetter + upperLetter + number)
	return stringRand(charset, length)
}

// StringRandNumber 随机数字
func StringRandNumber(length int) string {
	charset := []byte(number)
	return stringRand(charset, length)
}

// StringRandLetter 随机字母
func StringRandLetter(length int) string {
	charset := []byte(lowerLetter + upperLetter)
	return stringRand(charset, length)
}

// StringRandLowerLetter 随机小写字母
func StringRandLowerLetter(length int) string {
	charset := []byte(lowerLetter)
	return stringRand(charset, length)
}

// StringRandUpperLetter 随机大写字母
func StringRandUpperLetter(length int) string {
	charset := []byte(upperLetter)
	return stringRand(charset, length)
}

func stringRand(charset []byte, length int) string {
	//使用这种方式会导致for循环中随机值相同，直接使用rand即可
	//source := rand.NewSource(time.Now().UnixNano())
	//rng := rand.New(source)
	// 创建一个字节切片来保存随机字符
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))] // 从字符集中随机选择字符
	}
	return string(b)
}
