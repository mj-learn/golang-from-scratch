package main

import "fmt"

var myName string // not a problem if not used

func main() {
	fmt.Println("Hello, world.")

	var whatToSay string // problem if not used
	whatToSay = "Goodbye"
	fmt.Println(whatToSay)

	var i int
	i = 7
	fmt.Println("i is set to", i)

	whatWasSaid := saySomething()
	fmt.Println(whatWasSaid)

	firstThing, secoundThing := sayTwoThings()
	fmt.Println(firstThing, secoundThing)
}

func saySomething() string {
	return "something"
}

func sayTwoThings() (string, string) {
	return "something", "else"
}
