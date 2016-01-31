package helpers

import (
	"crypto/sha1"
	"encoding/hex"
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

func EncodeSha1(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
