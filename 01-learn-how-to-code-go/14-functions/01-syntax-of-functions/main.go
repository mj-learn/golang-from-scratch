package main

import "fmt"

// func (receiver) identifier(parameters) (returns) { code }

func main() {
	// foo
	foo()

	fmt.Println()

	// bar
	bar("MOnjofn")

	fmt.Println()

	// aloha
	s := aloha("MOn B. Jofn")
	fmt.Println(s)
	fmt.Println(aloha("Monjofn"))

	fmt.Println()

	// dogYears
	p1, p2 := dogYears("MOnjofn", 7)
	fmt.Println(p1, p2)
}

func foo() {
	fmt.Println("I'm from foo")
}

func bar(s string) {
	fmt.Println("My name is", s)
}

func aloha(s string) string {
	return fmt.Sprint("They call me Mr. ", s)
}

func dogYears(name string, age int) (string, int) {
	age *= 7 // short of age = age * 7
	return fmt.Sprint(name, " is this old in dog years"), age
}
