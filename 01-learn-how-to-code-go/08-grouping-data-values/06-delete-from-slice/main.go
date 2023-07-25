package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("xi = %#v\n", xi)
	xi = append(xi[:3], xi[5:7]...)
	fmt.Printf("xi = %#v\n", xi)
}
