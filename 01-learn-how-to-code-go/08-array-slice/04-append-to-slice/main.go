package main

import "fmt"

func main() {
	xi := []int{42, 43, 44}
	fmt.Println(xi)

	fmt.Println()

	xi = append(xi, 50)
	xi = append(xi, 60, 61, 62, 63)
	fmt.Println(xi)
}
