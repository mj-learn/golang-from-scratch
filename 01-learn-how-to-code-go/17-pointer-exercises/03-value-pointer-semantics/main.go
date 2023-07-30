package main

import "fmt"

type msg struct {
	first string
}

func main() {
	x := msg{"Helloo"}
	fmt.Println(x)
	changeOne(x)
	fmt.Println(x)
	x = changeOne(x)
	fmt.Println(x)

	fmt.Println()

	changeTwo(&x)
	fmt.Println(x)
}

func changeOne(m msg) msg {
	m.first = "Hello from changeOne"
	return m
}

func changeTwo(m *msg) {
	m.first = "Hello from changeTwo"
}

/*
Create two functions to change a field in a struct called "first" of TYPE string:
  ● One function will use VALUE SEMANTICS
      ○ this function will return some VALUE of some TYPE
  ● The other function will use POINTER SEMANTICS
      ○ this function will return nothing
  ● don't use methods
*/
