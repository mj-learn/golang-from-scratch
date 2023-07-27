package main

import "fmt"

func main() {
	x := func() {
		fmt.Println("Hello from x")
	}

	fmt.Printf("x -> %T\n", x)
}
