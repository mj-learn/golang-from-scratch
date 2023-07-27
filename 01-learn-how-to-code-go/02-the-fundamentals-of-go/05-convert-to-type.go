package main

import "fmt"

type age uint

func main() {
	a := 72
	b := age(a)
	c := string(b)
	// d := float64(c)
	// e := bool(c)

	fmt.Printf("(a) %T -> %v\n", a, a)
	fmt.Printf("(b) %T -> %v\n", b, b)
	fmt.Printf("(c) %T -> %v\n", c, c)
	// fmt.Printf("(d) %T -> %v\n", d, d)
	// fmt.Printf("(e) %T -> %v\n", e, e)

	fmt.Println()

	aa := "Hello world"
	bb := []byte(aa)

	fmt.Printf("(aa) %T -> %v\n", aa, aa)
	fmt.Printf("(bb) %T -> %v\n", bb, bb)
}
