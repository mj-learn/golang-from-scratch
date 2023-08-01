package main

import "fmt"

type myNumbers interface {
	int | string
}

func generic[T myNumbers](a, b T) {
	fmt.Println("Hello", a, b)
}

func main() {
	generic("2", "3")
	generic(1, 4)
}
