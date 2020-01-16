package go_util

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

func Sha256(str string) string {
	digest := sha256.Sum256([]byte(str))
	return hex.EncodeToString(digest[:])
}

func Md5(str string) string {
	digest := md5.Sum([]byte(str))
	return hex.EncodeToString(digest[:])
}


