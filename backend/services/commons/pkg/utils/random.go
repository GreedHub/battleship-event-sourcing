package utils

import (
	"math/rand"
	"time"
)

func getRandomString(symbols string, length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	for i := range b {
		b[i] = symbols[rand.Intn(len(symbols))]
	}
	return string(b)
}

func GetRandomString(length int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	return getRandomString(letters, length)
}

func GetRandomCapitalizedString(length int) string {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	return getRandomString(letters, length)
}
