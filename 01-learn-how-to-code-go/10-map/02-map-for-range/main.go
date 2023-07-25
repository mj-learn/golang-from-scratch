package main

import "fmt"

func main() {
	am := map[string]string{
		"task 1": "done",
		"task 2": "25% paused",
		"task 3": "done",
		"task 5": "80%",
		"task 6": "30%",
	}

	for key, val := range am {
		fmt.Println(key, val)
	}

	fmt.Println()

	for _, val := range am {
		fmt.Println(val)
	}
}
