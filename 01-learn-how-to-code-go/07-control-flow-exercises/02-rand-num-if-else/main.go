package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(251)

	fmt.Println("x =", x)
	fmt.Println("Value of x is:")

	if x >= 0 && x <= 100 {
		fmt.Println("between 0 and 100")
	} else if x > 100 && x <= 200 {
		fmt.Println("between 101 and 200")
	} else {
		fmt.Println("between 201 and 250")
	}
}
