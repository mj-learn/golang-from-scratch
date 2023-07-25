package main

import "fmt"

func main() {
	am := map[string]int{
		"John":    42,
		"Smith":   15,
		"MOn":     11,
		"MOnjofn": 105,
	}

	val, ok := am["Mon"]
	if ok {
		fmt.Println(val)
	} else {
		fmt.Println("Key didn't exist")
	}

	fmt.Println()

	if val, ok := am["MOnjofn"]; ok {
		fmt.Println(val)
		fmt.Println("value of ok is:", ok)
	} else {
		fmt.Println("Key didn't exist")
	}
}
