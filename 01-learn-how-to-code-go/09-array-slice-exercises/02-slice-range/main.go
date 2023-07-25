package main

import "fmt"

func main() {
	si := make([]int, 0, 10)
	si = append(si, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51)

	fmt.Println()
	fmt.Printf("si = %#v", si)
	fmt.Println()

	for key, val := range si {
		fmt.Printf("key: %v, val: %v, type: %T\n", key, val, val)
	}
}
