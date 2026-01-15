package util

import (
	"crypto/rand"
	"math/big"
	"strings"
)

// IsBlank checks whether the given string is blank.
func IsBlank(s string) bool {
	return strings.TrimSpace(s) == ""
}

const (
	lowerLetter = "abcdefghijklmnopqrstuvwxyz"
	upperLetter = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	number      = "0123456789"
)

// StringRand 随机字符串 (Secure)
func StringRand(length int) (string, error) {
	return randomString(lowerLetter+upperLetter+number, length)
}

// StringRandNumber 随机数字 (Secure)
func StringRandNumber(length int) (string, error) {
	return randomString(number, length)
}

// StringRandLetter 随机字母 (Secure)
func StringRandLetter(length int) (string, error) {
	return randomString(lowerLetter+upperLetter, length)
}

// StringRandLowerLetter 随机小写字母 (Secure)
func StringRandLowerLetter(length int) (string, error) {
	return randomString(lowerLetter, length)
}

// StringRandUpperLetter 随机大写字母 (Secure)
func StringRandUpperLetter(length int) (string, error) {
	return randomString(upperLetter, length)
}

func randomString(charset string, length int) (string, error) {
	if length <= 0 {
		return "", nil
	}
	b := make([]byte, length)
	maxIdx := big.NewInt(int64(len(charset)))
	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, maxIdx)
		if err != nil {
			return "", err
		}
		b[i] = charset[n.Int64()]
	}
	return string(b), nil
}
