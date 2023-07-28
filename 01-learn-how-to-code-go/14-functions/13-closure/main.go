package main

import "fmt"

func main() {
	f := do()

	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func do() func() int {
	a := 0
	return func() int {
		a += 1
		return a
	}
}
