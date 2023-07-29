package main

import "fmt"

func main() {
	a := appendTwo(5)
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
}

func appendTwo(n int) func() int {
	finalSum := n
	return func() int {
		finalSum += 2
		return finalSum
	}
}
