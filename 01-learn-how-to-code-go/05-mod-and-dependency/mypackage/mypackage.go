package mypackage

import (
  "05-mod-and-dependency/mypackage/mysubpackage"
  "fmt"
)

// Print "Hello from mypackage" in terminal
// Secound line
func PrintHello() {
  fmt.Println("Hello from mypackage")
  mysubpackage.HelloFromMySubPackage()
}
