package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

type square struct {
	length float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s square) area() float64 {
	return s.length * s.length
}

type calculator struct{}

func (c calculator) sumArea(shapes ...shape) float64 {
	var sum float64

	for _, shape := range shapes {
		sum += shape.area()
	}
	return sum
}

func main() {
	circle := circle{radius: 1}
	square := square{length: 1}
	cl := calculator{}
	sum := cl.sumArea(circle, square)
	fmt.Println("Sum of area is :", sum)
}
