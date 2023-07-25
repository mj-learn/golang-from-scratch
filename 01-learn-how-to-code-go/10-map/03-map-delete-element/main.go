package main

import "fmt"

func main() {
	am := make(map[int]int)
	am[1] = 134
	am[2] = 263
	am[3] = 22
	am[4] = 55
	am[5] = 98

	fmt.Println(am)

	delete(am, 2)
	// delete(am, 8) // won't panic

	fmt.Println(am)
	fmt.Println(am[55]) // won't panic
}
