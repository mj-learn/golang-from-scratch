package main

import "fmt"

type person struct {
	name string
}

type nickname struct {
	name string
}

type secretAgent struct {
	person
	nickname
	ltk bool
}

func (p person) whoiam() {
	fmt.Println("My name is", p.name)
}

func (p person) speak() {
	fmt.Println("I'm person", p.name)
}

func (agent secretAgent) speak() {
	fmt.Println("I'm secret agent", agent.person.name)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	sa := secretAgent{
		person:   person{name: "James Bond"},
		nickname: nickname{name: "007"},
		ltk:      true,
	}

	p := person{
		name: "John Smith",
	}

	sa.speak()         // speaks from secretAgent's methods
	sa.whoiam()        // secretAgent take whoiam func from person's methods
	sa.person.whoiam() // this works too
	sa.person.speak()  // secretAgent take speak funcfrom person's methods
	p.speak()

	fmt.Println()

	saySomething(sa)
	saySomething(p)
}
