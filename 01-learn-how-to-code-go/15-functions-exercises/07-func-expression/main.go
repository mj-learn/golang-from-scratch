package main

import "fmt"

func main() {
	a := func() string {
		return "Hello from func expression"
	}

	fmt.Println(a())
}
