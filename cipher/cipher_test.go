package cipher

import (
	"fmt"
	"testing"
)

func TestSha256Hex(t *testing.T) {
	str := "hello world"
	fmt.Println(Sha256Hex(str))
}

func TestHmacMd5Hex(t *testing.T) {
	str := "hello world"
	key := "123"
	fmt.Println(HmacMd5Hex(key, str))
}

func BenchmarkSha256Hex(b *testing.B) {
	str := "hello world"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sha256Hex(str)
	}
}
