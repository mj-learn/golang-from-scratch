package main

import "fmt"

func main() {
	x := 7
	fmt.Println(x)
	{
		fmt.Println(x)
		x := 5
		fmt.Println(x)
	}
}
