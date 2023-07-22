package main

import "fmt"

// Logical operators
// &&   AND
// ||   OR
// !    NOT
// true && true   TRUE
// true && false  FALSE
// true || true   TRUE
// true || false  TRUE
// !true          FALSE

func main() {
  fmt.Println("true && true = true ?")
  if true && true {
    fmt.Println("TRUE")
  }

  fmt.Println("")

  fmt.Println("true && false = false ?")
  if true && false {
    fmt.Println("TRUE")
  } else {
    fmt.Println("FALSE")
  }

  fmt.Println("")

  fmt.Println("true || true = true ?")
  if true || true {
    fmt.Println("TRUE")
  } else {
    fmt.Println("FALSE")
  }

  fmt.Println("")

  fmt.Println("true || false = true ?")
  if true || false {
    fmt.Println("TRUE")
  } else {
    fmt.Println("FALSE")
  }

  fmt.Println("")

  fmt.Println("!true = false ?")
  if !true {
    fmt.Println("TRUE")
  } else {
    fmt.Println("FALSE")
  }

  fmt.Println("")

  fmt.Println("true && !true = false ?")
  if true && !true {
    fmt.Println("TRUE")
  } else {
    fmt.Println("FALSE")
  }

  fmt.Println("")

  fmt.Println("!false && !false = true ?")
  if !false && !false {
    fmt.Println("TRUE")
  } else {
    fmt.Println("FALSE")
  }

}
