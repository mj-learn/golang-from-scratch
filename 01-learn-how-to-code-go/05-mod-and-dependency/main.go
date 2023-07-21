package main

import (
  "05-mod-and-dependency/mypackage"
  "05-mod-and-dependency/mypackage/mysubpackage"
  "05-mod-and-dependency/mypackage/mysubpackage/subsub"
  "fmt"

  "github.com/monjofn/puppy"
)


func main() {
  fmt.Println("Hello from main")
  puppy.Bark()
  puppy.BarkBark()
  mypackage.PrintHello()
  mysubpackage.HelloFromMySubPackage()
  subsub.HelloFromSubSub()
}
