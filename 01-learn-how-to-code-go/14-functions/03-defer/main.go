package main

import "fmt"

func main() {
	fmt.Println("top")
	boo()
	baz()
	defer foo()
	bar()
	too()
	fmt.Println("bot")
}

func foo() {
	fmt.Println("I'm from foo")
}

func bar() {
	fmt.Println("I'm from bar")
}

func baz() {
	fmt.Println("I'm from baz")
}

func too() {
	fmt.Println("I'm from too")
}

func boo() {
	fmt.Println("I'm from boo")
}
