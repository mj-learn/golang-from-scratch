package main

import "fmt"

func main() {
	xs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// fmt.Println(xs...)
	fmt.Println(sum(xs...))
}

func sum(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
