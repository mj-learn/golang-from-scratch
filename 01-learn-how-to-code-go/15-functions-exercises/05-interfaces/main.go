package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width  float64
}

func (s square) area() float64 {
	return s.width * s.length
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type shape interface {
	area() float64
}

func info(s shape) float64 {
	return s.area()
}

func main() {
	a := square{
		length: 5.5,
		width:  2,
	}

	b := circle{
		radius: 1,
	}

	fmt.Println(info(a))
	fmt.Println(info(b))
}
