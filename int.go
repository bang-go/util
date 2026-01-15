package util

import (
	"crypto/rand"
	"math/big"

	"github.com/bang-go/util/constraint"
)

// IntRandRange 区间随机数 (Secure)
func IntRandRange[T constraint.Integer](n1 T, n2 T) (int64, error) {
	v1, v2 := int64(n1), int64(n2)
	if v2 < v1 {
		v1, v2 = v2, v1
	}
	if v1 == v2 {
		return v1, nil
	}

	maxIdx := big.NewInt(v2 - v1)
	n, err := rand.Int(rand.Reader, maxIdx)
	if err != nil {
		return 0, err
	}
	return n.Int64() + v1, nil
}
