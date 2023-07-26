package main

import "fmt"

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	brand string
	make  int
	doors int
	color string
}

func main() {
	v1 := vehicle{
		engine: engine{electric: false},
		brand:  "Tesla",
		make:   1955,
		doors:  9,
		color:  "Pink",
	}

	v2 := vehicle{
		engine: engine{electric: true},
		brand:  "Lada",
		make:   2023,
		doors:  0,
		color:  "no",
	}

	cars := []vehicle{v1, v2}

	for _, val := range cars {
		fmt.Println("Brand:", val.brand)
		fmt.Println("Electric:", val.engine.electric)
		fmt.Println("Make:", val.make)
		fmt.Println("Doors:", val.doors)
		fmt.Println("Color:", val.color)
		fmt.Println()
	}
}
