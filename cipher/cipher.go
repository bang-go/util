package cipher

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

// Md5Hex returns the MD5 checksum of the data in hex string.
// Deprecated: MD5 is insecure. Use Sha256Hex instead.
func Md5Hex(s string) string {
	sum := md5.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}

// Sha1Hex returns the SHA1 checksum of the data in hex string.
// Deprecated: SHA1 is insecure. Use Sha256Hex instead.
func Sha1Hex(s string) string {
	sum := sha1.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}

// Sha256Hex returns the SHA256 checksum of the data in hex string.
func Sha256Hex(s string) string {
	sum := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sum[:])
}

// HmacMd5Hex returns the HMAC-MD5 of the data in hex string.
// Deprecated: MD5 is insecure. Use HmacSha256Hex instead.
func HmacMd5Hex(key, data string) string {
	hash := hmac.New(md5.New, []byte(key))
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}

// HmacSha256Hex returns the HMAC-SHA256 of the data in hex string.
func HmacSha256Hex(key, data string) string {
	hash := hmac.New(sha256.New, []byte(key))
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}

// HmacSha1Hex returns the HMAC-SHA1 of the data in hex string.
// Deprecated: SHA1 is insecure. Use HmacSha256Hex instead.
func HmacSha1Hex(key, data string) string {
	hash := hmac.New(sha1.New, []byte(key))
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}
