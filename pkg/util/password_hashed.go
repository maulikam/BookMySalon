package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const BcryptCostFactor = 12 // This should be the same value you use in your UserServiceImpl.

func main() {
	passwords := []string{"password", "password"}

	for _, password := range passwords {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), BcryptCostFactor)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(hashedPassword))
	}
}
