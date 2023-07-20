package main

import "fmt"

const (
  _ = iota // c0 == 0
  a
  b
  c
  d
  e
  f
)

const (
  aa = iota
  bb = iota + 4
  bb2
  cc = iota * 4
)

const (
  aaa = iota + 10
  bbb
  ccc
)

func main() {
  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)
  fmt.Println(e)
  fmt.Println(f)

  fmt.Println()

  fmt.Println(aa)
  fmt.Println(bb)
  fmt.Println(bb2)
  fmt.Println(cc)

  fmt.Println()

  fmt.Println(aaa)
  fmt.Println(bbb)
  fmt.Println(ccc)
  
  fmt.Println()

  fmt.Printf("%d \t %b\n", 1, 1)
  fmt.Printf("%d \t %b\n", 1<<a, 1<<a)
  fmt.Printf("%d \t %b\n", 1<<b, 1<<b)
  fmt.Printf("%d \t %b\n", 1<<c, 1<<c)
  fmt.Printf("%d \t %b\n", 1<<d, 1<<d)
  fmt.Printf("%d \t %b\n", 1<<e, 1<<e)
  fmt.Printf("%d \t %b\n", 1<<f, 1<<f)
}
