package main

import "fmt"

func main() {
	n := "MOnjofn"

	func() {
		fmt.Println("Hello from anonymous func")
	}()

	func(s string) {
		fmt.Println("Anonymous func with parameters by", s)
	}(n)
}
