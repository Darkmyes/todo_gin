package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	var passwordBytes = []byte(password)

	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(
		passwordBytes,
		bcrypt.DefaultCost,
	)

	if err != nil {
		fmt.Println(err)
	}

	return string(hashedPasswordBytes)
}

func DoPasswordsMatch(hashedPassword, requestPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(requestPassword))

	return err == nil
}

//https://www.gregorygaines.com/blog/posts/2020/11/21/how-to-properly-hash-and-salt-passwords-in-golang-bcrypt
