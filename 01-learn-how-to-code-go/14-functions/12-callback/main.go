package main

import "fmt"

/*
"callback" in programming refers to a function or a piece of code that is passed
as an argument to another function
*/

func main() {
	fmt.Printf("%T\n", sum)
	fmt.Printf("%T\n", subtract)
	fmt.Printf("%T\n", doMath)

	fmt.Println()

	x := doMath(1, 3, sum)
	fmt.Println(x)

	y := doMath(67, 21, subtract)
	fmt.Println(y)
}

func doMath(a, b int, f func(int, int) int) int {
	return f(a, b)
}

func sum(a, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}
