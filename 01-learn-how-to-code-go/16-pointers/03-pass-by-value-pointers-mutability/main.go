package main

import "fmt"

func intDelt(n *int) {
	*n += 1
}

// Becouse array, slice, map, channels are mutable
func sliceDelta(s []int) {
	s[0] = 555
}

func mapDelta(m map[string]int, s string) {
	m[s] = 555
}

func main() {
	x := 42
	fmt.Println(x)
	intDelt(&x)
	fmt.Println(x)

	fmt.Println()

	xs := []int{1, 2, 3, 4, 5}
	fmt.Println(xs)
	sliceDelta(xs)
	fmt.Println(xs)

	fmt.Println()

	xm := make(map[string]int)
	xm["James"] = 42
	fmt.Println(xm)
	mapDelta(xm, "James")
	fmt.Println(xm)
}
