package main

import "fmt"

type person struct {
	fistName string
	lastName string
	age      int
}

func main() {
	p1 := person{
		fistName: "MOn",
		lastName: "jofn",
		age:      15,
	}

	p2 := person{
		fistName: "James",
		lastName: "Bond",
		age:      42,
	}

	fmt.Printf("%T\n", p1)
	fmt.Printf("%#v\n", p1)
	fmt.Println("firstName is:", p1.fistName)
	fmt.Println("lastName is:", p1.lastName)
	fmt.Println("age is:", p1.age)

	fmt.Println()

	fmt.Println(p2)
}
