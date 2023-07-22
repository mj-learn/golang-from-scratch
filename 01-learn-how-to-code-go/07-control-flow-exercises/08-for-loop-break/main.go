package main

import "fmt"

func main() {
	x := 20
	for {
		if x == 5 {
			break
		}
		fmt.Println(x)
		x--
	}
}
