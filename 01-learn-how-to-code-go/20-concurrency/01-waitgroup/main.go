package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("ARCH:", runtime.GOARCH)
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	fmt.Println()

	wg.Add(1)
	go foo()
	bar()

	fmt.Println()

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	fmt.Println()

	wg.Wait()

	fmt.Println("end")
}

func foo() {
	for i := 0; i <= 5; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i <= 5; i++ {
		fmt.Println("bar:", i)
	}
}
