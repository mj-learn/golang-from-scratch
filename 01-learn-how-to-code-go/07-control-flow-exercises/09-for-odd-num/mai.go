package main

import "fmt"

func main() {
	for i := 50; i > 0; i-- {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
