package main

import "fmt"

func main() {
	//
	for i := 0; i < 5; i++ {
		fmt.Printf("%v ", i)
	}

	fmt.Println()

	y := 0
	for y < 5 {
		fmt.Printf("%v ", y)
		y++
	}

	fmt.Println()

	for {

		if y == 10 {
			break
		}

		if y == 8 {
			y++
			continue
		}

		fmt.Print("* ")

		y++
	}

	fmt.Println()

	x := 0
	for {
		if x == 3 {
			break
		}

		fmt.Println("Nested loop iteration: ", x)

		for i := 0; i <= 5; i++ {
			fmt.Printf("%v", i)
		}

		fmt.Println()

		x++
	}
}
