package main

import "fmt"

func main() {
	mds := make([][]string, 0, 2)
	mds = append(mds, []string{"James", "Bond", "Shaken, not stirred"},
		[]string{"Miss", "Moneypenny", "I'm 008"})
	fmt.Println(mds)

	for key, val := range mds {
		fmt.Printf("People %v: ", key+1)
		for _, val := range val {
			fmt.Printf("%v ", val)
		}
		fmt.Printf("\n")
	}
}
