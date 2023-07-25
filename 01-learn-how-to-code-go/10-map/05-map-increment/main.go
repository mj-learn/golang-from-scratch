package main

import "fmt"

func main() {
	am := map[string]int{
		"John":    42,
		"Smith":   15,
		"MOn":     11,
		"MOnjofn": 105,
	}
	fmt.Println(am)

	am["MOn"]++
	fmt.Println(am)

	am["MOn"]--
	fmt.Println(am)
}
