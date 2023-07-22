package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("This is where initialization for my program occurs")
}

func main() {
	x := rand.Intn(251)

	fmt.Println("x =", x)
	fmt.Println("Value of x is: ")

	switch {
	case x == 0, x <= 100:
		fmt.Println("between 0 and 100")
	case x > 100, x <= 200:
		fmt.Println("between 101 and 200")
	case x > 200, x <= 250:
		fmt.Println("between 201 and 250")
	}
}
