package main

import "fmt"

type person struct {
	first string
	age   int
}

func (p person) speak() {
	// fmt.Printf("Hello I'm %v and my age is %v\n", p.first, p.age)
	fmt.Println("Hello I'm", p.first, "and my age si", p.age)
}

func main() {
	p := person{
		first: "Mark",
		age:   5,
	}

	p.speak()
}
