package main

import "fmt"

func main() {
	si := []string{"a", "b", "c", "d"}
	fmt.Printf("si = %#v\n", si)
	fmt.Println()

	// Prefered method to create empty slice when you know how much elements will be
	xi := make([]int, 2, 10) // 2 is legth of slice, 10 is legth of array under slice
	fmt.Printf("xi = %#v\n", xi)
	fmt.Println("length is:", len(xi))
	fmt.Println("capacity is:", cap(xi))

	fmt.Println()
	xi = append(xi, 0, 1, 2, 3)
	fmt.Printf("xi = %#v\n", xi)
	fmt.Println("length is:", len(xi))
	fmt.Println("capacity is:", cap(xi))

	fmt.Println()
	xi = append(xi, 10, 11, 12, 13, 14, 15)
	fmt.Printf("xi = %#v\n", xi)
	fmt.Println("length is:", len(xi))
	fmt.Println("capacity is:", cap(xi))
}
