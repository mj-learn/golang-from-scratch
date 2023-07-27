package main

import (
	"fmt"
	"log"
)

type name struct {
	name string
}

type age int

func (n name) String() string {
	return fmt.Sprint("My name is ", n.name)
}

func (a age) String() string {
	return fmt.Sprint("My age is ", int(a))
}

// Wrapper func
func writeLog(s fmt.Stringer) {
	log.Println("LOG FROM (writeLog func)", s.String())
}

func main() {
	ns := name{
		name: "MOnjofn",
	}

	var a age = 234

	fmt.Println(ns)
	fmt.Println(a)

	writeLog(ns)
	writeLog(a)
}

/* This is build in golang
type Stringer interface {
    String() string
}
*/
