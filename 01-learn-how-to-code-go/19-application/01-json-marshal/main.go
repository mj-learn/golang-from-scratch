package main

import (
	"encoding/json"
	"fmt"
)

type people struct {
	Name string
	Age  int
}

func main() {
	p1 := people{Name: "Mon", Age: 11}
	p2 := people{Name: "Foo", Age: 22}

	ps, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(ps))

	peoples := []people{p1, p2}

	pss, err := json.Marshal(peoples)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(pss))
}
