package main

import "fmt"

func main() {
	x := a()
	fmt.Println(x())
}

func a() func() int {
	return b
}

func b() int {
	return 42
}
