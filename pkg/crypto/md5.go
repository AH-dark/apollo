package crypto

import (
	"crypto/md5"
	"fmt"
)

// MD5 returns the MD5 hash of the given data.
func MD5(data []byte) string {
	m := md5.New()
	m.Write(data)
	return fmt.Sprintf("%x", m.Sum(nil))
}

// MD5String returns the MD5 hash of the given string.
func MD5String(s string) string {
	return MD5([]byte(s))
}
