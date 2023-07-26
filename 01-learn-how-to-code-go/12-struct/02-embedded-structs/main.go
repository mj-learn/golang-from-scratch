package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

type secretAgent struct {
	person
	firstName string
	ltk       bool
}

func main() {
	sa1 := secretAgent{
		person: person{
			firstName: "James",
			lastName:  "Bond",
			age:       42,
		},
		firstName: "007",
		ltk:       true,
	}

	sa2 := secretAgent{
		person: person{
			firstName: "MOn",
			lastName:  "jofn",
			age:       15,
		},
		firstName: "ninja",
		ltk:       false,
	}

	fmt.Printf("%T\n", sa1)
	fmt.Printf("%#v\n", sa1)

	fmt.Println()
	fmt.Println("firstName is:", sa1.person.firstName)
	fmt.Println("lastName is:", sa1.person.lastName)
	fmt.Println("age is:", sa1.person.age)
	fmt.Println("firstName (nickname) is:", sa1.firstName)
	fmt.Println("ltk is:", sa1.ltk)

	fmt.Println()
	fmt.Println("secretAgent 2 is:", sa2)
}
