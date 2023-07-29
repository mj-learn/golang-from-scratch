package main

import "fmt"

func main() {
	foo()
	defer bar()
	defer baz()
	muu()
}

func foo() {
	fmt.Println("Hello from foo")
}

func bar() {
	fmt.Println("Hello from bar")
}

func baz() {
	fmt.Println("Hello from baz")
}

func muu() {
	fmt.Println("Hello from muu")
}
