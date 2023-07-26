package main

import "fmt"

func main() {
	as := struct {
		name      string
		friends   map[string]bool
		favDrinks []string
	}{
		name: "MOnjofn",
		friends: map[string]bool{
			"Elon Musk":                false,
			"Jerome Powell":            false,
			"Thomas A. Anderson (Neo)": true,
			"Satoshi Nakamoto":         true,
		},
		favDrinks: []string{"Gasoline", "Diesel"},
	}

	fmt.Println(as)
}
