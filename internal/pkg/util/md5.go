package util

import (
	"crypto/md5"
	"encoding/hex"
	"io"
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

func Md5Reader(reader io.Reader) (string, error) {
	h := md5.New()
	_, err := io.Copy(h, reader)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
