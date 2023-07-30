package main

import "fmt"

// pointer = memory adress (path to value in memory)

func main() {
	x := 42
	fmt.Println(x) // this print x

	// Pointers
	fmt.Printf("&x -> %T -> %v\n", &x, &x) // print type and path to memory

	fmt.Println()

	y := &x
	fmt.Printf("y -> %T -> %v\n", y, y) // print type and memory adress (path to value in memory)
	fmt.Println("Dereferencing pointer y -> *y ->", *y)

	*y = 11
	fmt.Println("x -> 42; *y = 11; x ->", x)
}
