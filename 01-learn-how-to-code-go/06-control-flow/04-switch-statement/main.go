package main

import "fmt"

func main() {
	x := 43
	y := 6

	switch {
	case x < 42:
		fmt.Println("x is less than 42")
	case x == 42:
		fmt.Println("x is equal to 42")
	case x > 42 && y == 5:
		fmt.Println("x is greater than 42 and y is 5")
	case x > 42:
		fmt.Println("x is greater than 42")
	}

	switch x {
	case 42:
		fmt.Println("x is 42")
	case 43:
		fmt.Println("x is 43")
	case 44:
		fmt.Println("x is 44")
	case 45:
		fmt.Println("x is 45")
	default:
		fmt.Println("x is something else")
	}

	switch y {
	case 5:
		fmt.Println("y is 5")
		fallthrough
	case 4:
		fmt.Println("y is 4")
	default:
		fmt.Println("default in y switch case")
	}
}
