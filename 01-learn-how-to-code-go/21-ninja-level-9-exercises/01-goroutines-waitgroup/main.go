package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Hello from main")

	wg.Add(1)
	go gr1()

	wg.Add(1)
	go gr2()

	wg.Wait()
}

func gr1() {
	fmt.Println("Go routine 1")
	wg.Done()
}

func gr2() {
	fmt.Println("Go routine 2")
	wg.Done()
}
