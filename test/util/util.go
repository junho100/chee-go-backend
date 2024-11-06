package util

import (
	"time"

	"golang.org/x/exp/rand"
)

func RandomString(n int) string {
	rand.Seed(uint64(time.Now().UnixNano()))
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
