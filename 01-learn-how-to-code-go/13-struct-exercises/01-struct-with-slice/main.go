package main

import "fmt"

type person struct {
	firsTname               string
	lastName                string
	favoriteIceCreamFlavors []string
}

func main() {
	p1 := person{
		firsTname:               "foo",
		lastName:                "bar",
		favoriteIceCreamFlavors: []string{"vanila", "chocolate"},
	}

	fmt.Printf("First name: %v\n", p1.firsTname)
	fmt.Printf("Last name: %v\n", p1.lastName)
	fmt.Print("Favorite Ice Cream Flavors: ")

	for _, val := range p1.favoriteIceCreamFlavors {
		fmt.Printf("%v, ", val)
	}

	fmt.Println()
}
