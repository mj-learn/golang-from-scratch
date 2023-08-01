package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type myAlias int

type myNymbers interface {
	constraints.Integer | constraints.Float | constraints.Signed
}

func generic[T myNymbers](a, b T) {
	fmt.Println(a, b)
}

func main() {
	var x myAlias = 55
	var y myAlias = 3
	generic(1, 2)
	generic(1.4, 2.55)
	generic(x, y)
}
