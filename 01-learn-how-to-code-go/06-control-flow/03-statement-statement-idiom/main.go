package main

import (
  "fmt"
  "math/rand"
)

func main() {

  x := 5; y := x + 5
  fmt.Printf("y is %v + 5 = %v\n", x, y)

  fmt.Println()
  
  // Statement, statement ideom
  // if z:=5; foo := z * rand.Intin(z); foo >= z  -> this will get error

  if foo := x * rand.Intn(x); foo >= x {
    fmt.Printf("foo is %v and is GREATER OR EQUAL to x which is %v", foo, x)
  } else {
    fmt.Printf("foo is %v and is LESS THAN x which is %v", foo, x)
  }

  fmt.Println()
}
