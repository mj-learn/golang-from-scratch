package main

import "fmt"

func main() {
	x := "MOn"
	fmt.Println("x := \"MOn\"")
	fmt.Println("x ->", x)     // value
	fmt.Println("&x ->", &x)   // pointer -> memory adress (path to value in memory)
	fmt.Println("*&x ->", *&x) // dereferencing (get the value from memory adress)

	fmt.Println()
	fmt.Println("y := &x")

	y := &x
	fmt.Println("y ->", y)   // pointer to x
	fmt.Println("*y ->", *y) // dereferencing (get the value from memory adress)

	fmt.Println()
	fmt.Println("*y = \"MOnjofn\"")

	*y = "MOnjofn"
	fmt.Println("y ->", y)   // value
	fmt.Println("*y ->", *y) // dereferencing (get the value from memory adress)
	fmt.Println("x ->", x)   // value
}
