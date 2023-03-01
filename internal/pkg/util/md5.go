package util

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	cipherStr := h.Sum(nil)

	return hex.EncodeToString(cipherStr)
}

func Md5Equal(secret, plain string) bool {
	return Md5(plain) == secret
}
