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
	fmt.Printf("%#v\n", am)
	fmt.Println("MOn`s age is:", am["MOn"])

	fmt.Println()

	bm := make(map[string]int)
	fmt.Println(bm)
	fmt.Printf("%#v\n", bm)

	fmt.Println()

	bm["foo"] = 16
	bm["aaa"] = 11
	bm["bar"] = 17
	bm["baz"] = 71
	fmt.Println(bm)
	fmt.Printf("%#v\n", bm)

	fmt.Println()
	fmt.Println("Length of bm map is:", len(bm))
}
