package main

import "fmt"

type myAge int

type myNumbers interface {
	~int | ~float64
}

func generic[T myNumbers](a, b T) {
	fmt.Println(a, b)
}

func genericTwo[T ~int | string](a, b T) {
	fmt.Println(a, b)
}

func main() {
	x := myAge(21)
	var y myAge = 42

	generic(x, y)
	generic(1, 2)

	genericTwo(x, y)
	genericTwo(5, 12)
}
