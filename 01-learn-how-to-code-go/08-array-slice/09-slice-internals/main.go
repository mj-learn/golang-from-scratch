package main

import "fmt"

func main() {
	a := "MOn"
	b := a
	fmt.Println("a =", a, "\t b =", b)
	a = "MOnjofn"
	fmt.Println("a =", a, "\t b =", b)

	fmt.Println()

	x := 5
	y := x
	fmt.Println("x =", x, "\t y =", y)
	x = 7
	fmt.Println("x =", x, "\t y =", y)

	fmt.Println()

	xi := []int{1, 2, 3}
	yi := xi
	fmt.Printf("xi = %#v\n", xi)
	fmt.Printf("yi = %#v\n", yi)

	fmt.Println()

	xi[0] = 7
	fmt.Printf("xi = %#v\n", xi)
	fmt.Printf("yi = %#v\n", yi)
}
