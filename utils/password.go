package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashingPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func VerifyPassword(hashedPassword string, password string) bool {
	fmt.Println("Hashed Password:", hashedPassword)
    fmt.Println("Password:", password)
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	fmt.Println(err)
	return err == nil
}