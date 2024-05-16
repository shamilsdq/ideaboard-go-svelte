package utils

import "math/rand"

func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz"
	const charsetLen = len(charset)

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Int63()%int64(charsetLen)]
	}
	return string(b)
}
