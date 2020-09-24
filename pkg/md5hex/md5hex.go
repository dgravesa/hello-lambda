package md5hex

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

// Sum computes a hex-encoded representation of a string's MD5 sum
func Sum(s string) string {
	bytes := []byte(s)
	sumBytes := md5.Sum(bytes)
	return hex.EncodeToString(sumBytes[:])
}

// ReadSum computes a hex-encoded representation of a stream's MD5 sum
func ReadSum(r io.Reader) (string, error) {
	h := md5.New()
	if _, err := io.Copy(h, r); err != nil {
		return "", err
	}
	sumBytes := h.Sum(nil)
	return hex.EncodeToString(sumBytes), nil
}
