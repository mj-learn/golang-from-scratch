package main

import (
	"fmt"
	"sort"
)

func main() {
	s1 := []int{10, 5, 22, 1, 11}
	s2 := []string{"q", "Q", "Mon", "MOn", "jofn", "JOFN", "J0fn"}

	fmt.Println(s1)
	sort.Ints(s1)
	fmt.Println(s1)

	fmt.Println()

	fmt.Println(s2)
	sort.Strings(s2)
	fmt.Println(s2)
}
