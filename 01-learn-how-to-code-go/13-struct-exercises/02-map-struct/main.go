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

	p2 := person{
		firsTname:               "MOn",
		lastName:                "jofn",
		favoriteIceCreamFlavors: []string{"pizza", "chips"},
	}

	peoples := make(map[string]person)
	peoples["p1"] = p1
	peoples["p2"] = p2

	for _, p := range peoples {
		fmt.Printf("First name: %v\n", p.firsTname)
		fmt.Printf("Last name: %v\n", p.lastName)
		fmt.Print("Ice Cream Favors: ")
		for _, favor := range p.favoriteIceCreamFlavors {
			fmt.Printf("%v, ", favor)
		}
		fmt.Println()
		fmt.Println()
	}
}
