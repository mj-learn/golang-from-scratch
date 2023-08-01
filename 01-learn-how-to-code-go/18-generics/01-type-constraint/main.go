package main

import "fmt"

func generic[T int | string](a, b T) {
	fmt.Println("Hello", a, b)
}

func main() {
	generic("2", "3")
	generic(1, 4)
}
