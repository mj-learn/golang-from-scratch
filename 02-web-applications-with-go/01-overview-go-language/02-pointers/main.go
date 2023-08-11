package main

import "log"

func main() {
	myString := "Hi"
	log.Println("myString ->", myString)
	changeWithoutPointer(myString)
	log.Println("myString after changeWithoutPointer ->", myString)
	changeUsingPointer(&myString)
	log.Println("myString after changeUsingPointer ->", myString)
}

func changeWithoutPointer(s string) {
	newString := "Hello, world"
	s = newString
	log.Println("myString in changeWithoutPointer scope ->", s)
}

func changeUsingPointer(s *string) {
	newString := "Hello"
	*s = newString
}
