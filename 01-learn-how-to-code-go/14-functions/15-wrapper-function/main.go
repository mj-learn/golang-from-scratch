package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	xb, err := readFile("./15-wrapper-function/poem.txt")
	if err != nil {
		log.Fatalf("Error in main in readFile: %s", err)
	}
	fmt.Println(string(xb))
}

func readFile(fileName string) ([]byte, error) {
	xb, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("Error in readFile: %s", err)
	}
	return xb, nil
}
