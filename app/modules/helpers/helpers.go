package helpers

import (
	"math/rand"
	"time"
)

var (
	alphabet = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
)

func RandomString(n int) string {
	rand.Seed(time.Now().UTC().UnixNano())

	result := make([]rune, n)
	for i := range result {
		result[i] = alphabet[rand.Intn(len(alphabet))]
	}

	return string(result)
}
