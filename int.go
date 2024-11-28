package util

import (
	"github.com/bang-go/util/constraint"
	"math/rand"
	"strconv"
)

// IntToString Int转化String(10进制)
func IntToString[T constraint.Integer](val T) (str string) {
	str = strconv.FormatInt(int64(val), 10)
	return
}

// IntRandRange 区间随机数
func IntRandRange[T constraint.Integer](n1 T, n2 T) int64 {
	if n2 < n1 {
		n1, n2 = n2, n1
	}
	if n1 == n2 {
		return int64(n1)
	}
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return rand.Int63n(int64(n2-n1)) + int64(n1)
}
