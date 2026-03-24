package cipher

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

// SHA256Hex returns the SHA256 checksum of the data in hex string.
func SHA256Hex(s string) string {
	sum := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sum[:])
}

// HMACSHA256Hex returns the HMAC-SHA256 of the data in hex string.
func HMACSHA256Hex(key, data string) string {
	hash := hmac.New(sha256.New, []byte(key))
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}

// HMACSHA1Hex returns the HMAC-SHA1 of the data in hex string.
// Deprecated: SHA-1 is insecure. Prefer HMACSHA256Hex for new code.
func HMACSHA1Hex(key, data string) string {
	hash := hmac.New(sha1.New, []byte(key))
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}

// HmacSha1Hex returns the HMAC-SHA1 of the data in hex string.
// Deprecated: use HMACSHA1Hex.
func HmacSha1Hex(key, data string) string {
	return HMACSHA1Hex(key, data)
}
