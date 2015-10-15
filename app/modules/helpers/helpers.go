package helpers

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
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

func JoinStrings(strings ...string) string {
	var buf bytes.Buffer

	for _, s := range strings {
		buf.WriteString(s)
	}

	return buf.String()
}

func EncodeSha1(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
