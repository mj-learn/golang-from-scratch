package main

import "fmt"

func main() {
	nums(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println()

	fmt.Println(sum("Just to test it", 1, 2, 3, 4, 5, 6, 7, 8, 9))
}

func nums(num ...int) {
	fmt.Println(num)
	fmt.Printf("%T\n", num)
}

// som of the int elements
func sum(s string, num ...int) (string, []int, int) {
	sum := 0
	for _, val := range num {
		sum += val
	}
	return s, num, sum
}
