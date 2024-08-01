package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string,error) {
	// Hashing the password with the cost of 14
	hashed,err :=bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "",err
	}
	return string(hashed),nil
}

func ComparePassword(hashedPassword string, password string) bool {
	fmt.Println(hashedPassword)
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}