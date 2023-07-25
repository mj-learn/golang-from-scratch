package main

import (
	"fmt"
	"sort"
)

func sortOne(s []int) []int {
	sort.Ints(s)
	return s
}

func sortTwo(s []int) []int {
	n := make([]int, len(s))
	copy(n, s)
	sort.Ints(n)
	return n
}

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

	xi := []int{5, 3, 1}
	yi := xi             // create pointer to underlying array
	ci := make([]int, 3) // create new slice
	ci[0] = 8
	fmt.Printf("ci = %#v\n", ci)
	copy(ci, xi) // copy elements from xi in ci
	fmt.Printf("xi = %#v\n", xi)
	fmt.Printf("yi = %#v\n", yi)
	fmt.Printf("ci = %#v\n", ci)

	fmt.Println()

	xi[0] = 7
	fmt.Printf("xi = %#v\n", xi)
	fmt.Printf("yi = %#v\n", yi)
	fmt.Printf("ci = %#v\n", ci)

	fmt.Println()

	fmt.Println("xi (before sortOne) -", xi)
	fmt.Println("sortOne return:", sortOne(xi))
	fmt.Println("xi (after sortOne) -", xi)

	fmt.Println()

	fmt.Println("ci (before sortOne) -", ci)
	fmt.Println("sortTwo return:", sortTwo(ci))
	fmt.Println("ci (after sortOne) -", ci)
}
