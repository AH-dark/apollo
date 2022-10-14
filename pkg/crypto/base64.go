package crypto

import "encoding/base64"

// Base64Encode encodes a string to base64
func Base64Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// Base64Decode decodes a base64 string
func Base64Decode(s string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(s)
	return string(decoded), err
}
