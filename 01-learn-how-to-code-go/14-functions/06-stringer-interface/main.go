package main

import "fmt"

type name struct {
	name string
}

type job struct {
	name string
}

type age int

func (n name) String() string {
	return fmt.Sprint("My name is ", n.name)
}

func (j job) String() string {
	return fmt.Sprint("My job is ", j.name)
}

func (a age) String() string {
	return fmt.Sprint("My age is ", int(a))
}

func main() {
	ns := name{
		name: "MOnjofn",
	}

	js := job{
		name: "Clown",
	}

	age := 2005

	fmt.Println(ns)
	fmt.Println(js)
	fmt.Println(age)
}

/* This is build in golang
type Stringer interface {
    String() string
}
*/
