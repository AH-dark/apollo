package crypto

func Password(pass string) string {
	return string(MD5String(pass))
}
