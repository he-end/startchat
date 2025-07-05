package authpassword

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"log"

	"golang.org/x/crypto/argon2"
)

type Params struct {
	Memory      uint32
	SaltLength  uint32
	Iterations  uint32
	Parallelism uint8
	KeyLength   uint32
}

var p = Params{
	Memory:      64 * 124,
	SaltLength:  16,
	Iterations:  3,
	KeyLength:   32,
	Parallelism: 2,
}

/* result is base64 encoding*/
func HashingPassword(pwd string) (pwdHashed string, err error) {
	hash, salt, err := generateHash(pwd, &p)
	if err != nil {
		log.Println("error : ", err)
		return "", err
	}

	base64Hash := base64.StdEncoding.EncodeToString(hash)
	base64Salt := base64.StdEncoding.EncodeToString([]byte(salt))
	finalHash := base64Salt + "$" + base64Hash
	return finalHash, nil
}

func generateRandBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, errors.New("cant read random byte")
	}
	return b, nil
}

func generateHash(pwd string, p *Params) (hash []byte, salt []byte, err error) {
	salt, err = generateRandBytes(p.SaltLength)
	if err != nil {
		log.Fatal("error get random byte")
		return nil, nil, err
	}
	// hash := argon2.IDKey([]byte(pwd), salt, p.iterations, p.memory, p.parallelism, p.keyLength)
	hash = argon2.IDKey([]byte(pwd), salt, p.Iterations, p.Memory, p.Parallelism, p.KeyLength)
	return hash, salt, nil

}
