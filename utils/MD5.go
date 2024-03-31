package utils

import (
	"crypto/md5"
	"fmt"
)

func MD5(data []byte) string {
	hash := md5.New()
	hash.Write(data)

	hashed := hash.Sum(nil)
	return fmt.Sprintf("%x", hashed)
}
