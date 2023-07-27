package main

import "fmt"

func main() {
	x := foo()
	y := bar()
	z := baz()

	fmt.Printf("x -> %T\n", x)
	fmt.Printf("y -> %T\n", y)
	fmt.Printf("z -> %T\n", z)
	fmt.Printf("y() -> %T -> %v\n", y(), y())
	fmt.Printf("z() -> %T -> %v\n", z(), z())
}

func foo() int {
	return 42
}

func bar() func() int {
	return foo
}

func baz() func() string {
	return func() string {
		return "Hello from baz sub func"
	}
}
