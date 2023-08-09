package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	p := "password@123"
	hp, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
	if err != nil {
		fmt.Println("YOU CAN'T LOGIN")
	}
	fmt.Println(string(hp))
}
