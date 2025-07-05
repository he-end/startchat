package authrandtoken

import (
	"crypto/rand"
	"math/big"
)

var KeyRandT string

func init() {
	KeyRandT, _ = generateRandomKey(40)
}
func generateRandomKey(length int) (string, error) {
	comb := []string{
		"abcdefghijklmnopqrstuvwzyz",
		"ABCDEFGHIJKLMNOPQRSTUVEXYZ",
		"0123456789",
		"!@#$%^&^&*()_+?/:;{},.<>",
	}

	// Gabungin semua karakter jadi satu string
	charset := ""
	for _, c := range comb {
		charset += c
	}

	result := make([]byte, length)
	for i := 0; i < length; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result[i] = charset[index.Int64()]
	}
	return string(result), nil
}
