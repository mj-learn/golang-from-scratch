package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	for key, val := range arr {
		fmt.Printf("key: %v, val: %v, type:%T\n", key, val, val)
	}
}
