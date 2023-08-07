package main

import (
	"encoding/json"
	"fmt"
)

// use https://mholt.github.io/json-to-go/ to generate
type person struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

func main() {
	s := `[{"Name":"Mon","Age":11},{"Name":"Foo","Age":22}]`
	bs := []byte(s)
	peoples := []person{}

	err := json.Unmarshal(bs, &peoples)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(peoples)

	for _, v := range peoples {
		fmt.Println(v.Age, v.Name)
	}
}
