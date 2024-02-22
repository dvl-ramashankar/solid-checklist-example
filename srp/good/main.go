package main

import (
	"fmt"
	"math"
)

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

func printArea(value float64, tp string) {
	fmt.Printf("Area of the %s is: %f", tp, value)
}

func main() {
	circle := circle{radius: 1}
	printArea(circle.area(), "circle")
	fmt.Println()
	square := square{length: 1}
	printArea(square.area(), "square")
}
