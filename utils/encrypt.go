package utils

import (
	"crypto/rand"
	"encoding/base64"
	"golang.org/x/crypto/scrypt"
	"log"
	"math/big"
)

const KeyLen = 10

// ScryptPw 密码加密
func ScryptPw(password string) (string, int) {
	salt, _ := rand.Int(rand.Reader, big.NewInt(20))
	hashPw, err := scrypt.Key([]byte(password), salt.Bytes(), 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(hashPw)
	return fpw, salt.Sign()
}
