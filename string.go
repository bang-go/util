package util

import (
	"crypto/rand"
	"errors"
	"math/big"
	"strings"
)

// ErrNegativeLength indicates that a helper received a negative length.
var ErrNegativeLength = errors.New("util: negative length")

const (
	lowerLetters = "abcdefghijklmnopqrstuvwxyz"
	upperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits       = "0123456789"
)

// IsBlank checks whether the given string is blank.
func IsBlank(s string) bool {
	return strings.TrimSpace(s) == ""
}

// StringRand returns a cryptographically secure random alphanumeric string.
func StringRand(length int) (string, error) {
	return randomString(lowerLetters+upperLetters+digits, length)
}

// StringRandNumber returns a cryptographically secure random numeric string.
func StringRandNumber(length int) (string, error) {
	return randomString(digits, length)
}

// StringRandLetter returns a cryptographically secure random alphabetic string.
func StringRandLetter(length int) (string, error) {
	return randomString(lowerLetters+upperLetters, length)
}

// StringRandLowerLetter returns a cryptographically secure random lowercase string.
func StringRandLowerLetter(length int) (string, error) {
	return randomString(lowerLetters, length)
}

// StringRandUpperLetter returns a cryptographically secure random uppercase string.
func StringRandUpperLetter(length int) (string, error) {
	return randomString(upperLetters, length)
}

func randomString(charset string, length int) (string, error) {
	if length < 0 {
		return "", ErrNegativeLength
	}
	if length == 0 {
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
