package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("xi - %#v\n", xi)
	fmt.Println("")

	// [inclusive:exclusive]
	fmt.Printf("xi[2:4] - %#v\n\n", xi[2:4])

	// [:exclusive]
	fmt.Printf("xi[:4] - %#v\n\n", xi[:4])

	// [inclusive:]
	fmt.Printf("xi[2:] - %#v\n\n", xi[2:])

	//[:]
	fmt.Printf("xi[:] - %#v\n\n", xi[:])

	for key, val := range xi[:4] {
		fmt.Println(key, val)
	}

	fmt.Println()

	xi = xi[:4]
	fmt.Printf("xi = %#v\n", xi)
}
