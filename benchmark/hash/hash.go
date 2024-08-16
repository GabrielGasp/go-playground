package hash

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

func Sum256(s string) string {
	hash := sha256.Sum256([]byte(s))
	return hex.EncodeToString(hash[:])
}

func SumMD5(s string) string {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}
