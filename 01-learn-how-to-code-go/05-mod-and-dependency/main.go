package main

import (
	"fmt"

	"github.com/monjofn/puppy"

	"05-mod-and-dependency/mypackage"
	"05-mod-and-dependency/mypackage/mysubpackage"
	"05-mod-and-dependency/mypackage/mysubpackage/subsub"
)

func main() {
	fmt.Println("Hello from main")
	puppy.Bark()
	puppy.BarkBark()
	mypackage.PrintHello()
	mysubpackage.HelloFromMySubPackage()
	subsub.HelloFromSubSub()
}
