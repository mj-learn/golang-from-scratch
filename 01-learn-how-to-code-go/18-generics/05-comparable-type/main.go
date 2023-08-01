package main

import "fmt"

func sum[K comparable, V int64 | float64](a map[K]V) V {
	var s V
	for _, v := range a {
		s += v
	}
	return s
}

func main() {
	xm := map[string]int64{
		"one":  1,
		"two":  2,
		"tree": 3,
	}

	ym := make(map[int]float64)
	ym[1] = 1.5
	ym[2] = 2.5
	ym[3] = 3.5

	fmt.Println(sum(xm))
	fmt.Println(sum(ym))

	fmt.Println()

	fmt.Printf("%T\n", sum(xm))
	fmt.Printf("%T\n", sum(ym))
}
