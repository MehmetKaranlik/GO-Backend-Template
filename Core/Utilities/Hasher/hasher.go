package Hasher

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) []byte {

	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MaxCost)
	if err != nil {
		fmt.Printf("Error while encrypting token: %s", err.Error())
	}
	return bs
}

func ComparePassword(password string, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err != nil {
		fmt.Printf("Error while comparing token: %s", err.Error())
		return false
	}
	return true
}
