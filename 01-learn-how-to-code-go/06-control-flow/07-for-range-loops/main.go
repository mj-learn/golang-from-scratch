package main

import "fmt"

func main() {
	fmt.Println("In \"MOnjofn\": ")
	for key, val := range "MOnjofn" {
		fmt.Printf("Key is %v and val is %v and symbol is %#U\n", key, val, val)
	}

	fmt.Println()

	// for range loop
	// data structures - slice
	xi := []int{42, 43, 44, 45, 46, 47}
	for key, val := range xi {
		fmt.Println("ranging over a slice", key, val)
	}

	fmt.Println()

	// for range loop
	// data structures - map
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}
	for key, val := range m {
		fmt.Println("ranging over a map", key, val)
	}

	fmt.Println()

	// use only value
	for _, val := range xi {
		fmt.Println("ranging over a slice", val)
	}

	fmt.Println()

	// use only key
	for key := range m {
		fmt.Println("ranging over a map", key)
	}
}
