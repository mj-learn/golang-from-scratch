package main

import "fmt"

const Pi = 3.14

func main() {
  // Constants cannot be declared using the := syntax.

  const World = "World"
  fmt.Println("Hello", World)
  fmt.Println("Happy", Pi, "Day")

  const Truth = true
  fmt.Println("Go rules?", Truth)
}
