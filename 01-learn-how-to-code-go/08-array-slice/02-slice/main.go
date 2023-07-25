package main

import "fmt"

func main() {
	a := []int{}
	b := []string{"M", "MOn", "MOnjofn"}

	fmt.Println(a, b) // print slices
	fmt.Println(b[1]) // print MOn
	fmt.Printf("b type is %T\n", b)
}
