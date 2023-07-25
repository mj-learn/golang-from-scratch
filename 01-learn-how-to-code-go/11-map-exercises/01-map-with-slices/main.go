package main

import "fmt"

func main() {
	ms := make(map[string][]string)
	ms["bond_james"] = append(ms["bond_james"], "shaken, not stirred", "martinis", "fast cars") // just to test it
	ms["moneypenny_jenny"] = []string{"intelligence", "literature", "computer science"}
	ms["no_dr"] = []string{"cats", "ice cream", "sunsets"}

	for k, v := range ms {
		fmt.Printf("%v: ", k)
		for _, v := range v {
			fmt.Printf("%v, ", v)
		}
		fmt.Print("\n")
	}

	delete(ms, "no_dr")
	fmt.Println()

	for k, v := range ms {
		fmt.Printf("%v: ", k)
		for _, v := range v {
			fmt.Printf("%v, ", v)
		}
		fmt.Print("\n")
	}
}
