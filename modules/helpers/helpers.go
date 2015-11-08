package helpers

import (
	"math/rand"
	"time"
)

var (
	alphabet = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
)

// RandomString generates random string with given len.
func RandomString(n int) string {
	rand.Seed(time.Now().UTC().UnixNano())

	result := make([]rune, n)
	for i := range result {
		result[i] = alphabet[rand.Intn(len(alphabet))]
	}

	return string(result)
}
