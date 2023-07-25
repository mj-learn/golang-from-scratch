package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Martini", "Chocolate"}
	jm := []string{"Jenny", "Moneypenny", "Vodka", "Vanila"}

	// Multidimensional slice is slice of slices
	xxs := [][]string{jb, jm, {"MOn", "MOnjofn", "Water", "Oreo"}}
	fmt.Println(xxs)
}
