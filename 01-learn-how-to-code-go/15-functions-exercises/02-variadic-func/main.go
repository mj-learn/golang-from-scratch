package main

import "fmt"

func main() {
	xs := []int{1, 5, 4}
	fmt.Println(foo(xs...)) // with variadic
	fmt.Println(bar(xs))    // without variadic
}

func foo(i ...int) int {
	sum := 0
	for _, v := range i {
		sum += v
	}
	return sum
}

func bar(i []int) int {
	s := 0
	for _, v := range i {
		s += v
	}
	return s
}
