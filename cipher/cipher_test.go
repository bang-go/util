package cipher

import (
	"testing"
)

func TestHexDigests(t *testing.T) {
	input := "hello world"

	if got := SHA256Hex(input); got != "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9" {
		t.Fatalf("SHA256Hex() = %q", got)
	}
}

func TestHMACDigests(t *testing.T) {
	key := "123"
	input := "hello world"

	if got := HMACSHA256Hex(key, input); got != "8de9bbe5596700556793559dd70e6486684f379d162f70a4a341b3ee383565d4" {
		t.Fatalf("HMACSHA256Hex() = %q", got)
	}
}

func BenchmarkSHA256Hex(b *testing.B) {
	str := "hello world"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SHA256Hex(str)
	}
}
