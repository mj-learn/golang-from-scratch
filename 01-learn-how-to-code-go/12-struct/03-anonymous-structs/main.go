package main

import "fmt"

func main() {
	as := struct {
		a bool
		b float32
		c int
	}{
		a: true,
		b: 3.14,
		c: 42,
	}

	fmt.Printf("%T\n", as)
	fmt.Printf("%#v\n", as)
	fmt.Printf("%v\n", as)

	fmt.Println()

	fmt.Println("a is:", as.a)
	fmt.Println("b is:", as.b)
	fmt.Println("c is:", as.c)
}
