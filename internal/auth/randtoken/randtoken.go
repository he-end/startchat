package authrandtoken

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
)

func GenerateSecureRandomToken() (string, error) {
	token := make([]byte, 32)
	_, err := rand.Read(token)
	if err != nil {
		return "", errors.New("error create secure token")
	}
	return base64.StdEncoding.EncodeToString(token), nil
}

func HashRanomToken(token string, secret_key string) string {
	hash := hmac.New(sha256.New, []byte(secret_key))
	hash.Write([]byte(token))
	return hex.EncodeToString(hash.Sum(nil))
}
