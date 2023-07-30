package main

import "fmt"

type dog struct {
	name string
}

func (d dog) walk() {
	fmt.Println("My name is", d.name, "and I'm walking.")
}

func (d *dog) run() {
	d.name = "Foo"
	fmt.Println("My name is", d.name, "and I'm running.")
}

type younin interface {
	walk()
	run()
}

func youngRun(y younin) {
	y.run()
}

func youngWalk(y younin) {
	y.walk()
}

func main() {
	d1 := dog{"Bar"}
	d1.walk()
	d1.run()
	// youngRun(d1)
	// youngWalk(d1)

	fmt.Println()

	d2 := &dog{"Pointer"}
	d2.walk()
	d2.run()
	youngRun(d2)
	youngWalk(d2)
}
