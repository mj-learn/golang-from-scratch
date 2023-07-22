package main

import "fmt"

// ==   equal
// !=   not equal
// <    less
// <=   less or equal
// >    greater
// >=   greater or equal

func main() {

  if true {
    fmt.Println("true")
  }

  if !false {
    fmt.Println("true")
  }

  x, y := 42, 40

  if x == 42 {
    fmt.Println("x is 42")
  }

  if x != 42 {
    fmt.Println("x is not 42")
  }

  if y > 42 {
    fmt.Println("y is greater than 42")
  } else if y == 42 {
    fmt.Println("y is 42")
  } else {
    fmt.Println("y is less than 42")
  }

}
