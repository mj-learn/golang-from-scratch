package main

import "fmt"

// A method is nothing more than a FUNC attached to a TYPE

/*
The type they are associated with is called the receiver type, and the method
can be called using any variable of that type
*/

// syntax: func (receiver) identifier(parameters) (returns) { code }

type person struct {
	name string
}

func (freak person) speak() {
	fmt.Println("Hi :) I'm", freak.name, "from Mars")
}

func main() {
	alien := person{
		name: "Elon",
	}

	alien.speak()
}
