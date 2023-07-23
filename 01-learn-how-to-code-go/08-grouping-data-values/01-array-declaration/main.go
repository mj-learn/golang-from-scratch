package main

import "fmt"

func main() {
	var a [5]int // array declaration

	b := [4]int{} // short array declaration
	b[1] = 20

	fmt.Println("array a:", a)
	fmt.Println("array b:", b)

	fmt.Println()

	c := [3]int{42, 43, 44}
	fmt.Println("array c:", c)
	fmt.Printf("array c (from printf): %v\n", c)

	fmt.Println()

	d, f := [2]int{}, [2]int{} // declaration of multiple arrays at once
	fmt.Println("array d:", d)
	fmt.Println("array f:", f)

	fmt.Println()

	x := [...]string{"M", "MJ", "MOn", "MOnjofn"}
	fmt.Println(x)

	fmt.Println()

	fmt.Printf("x is %T\n", x)
	fmt.Println("Len of x is:", len(x))
}
