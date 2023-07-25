package main

import "fmt"

func main() {
	si := make([]int, 0, 10)
	si = append(si, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51)

	si_1 := si[:5]
	si_2 := si[5:]
	si_3 := si[2:7]
	si_4 := si[1:6]

	fmt.Println(si_1)
	fmt.Println(si_2)
	fmt.Println(si_3)
	fmt.Println(si_4)
}
