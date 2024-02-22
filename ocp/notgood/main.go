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

type calculator struct{}

func (c calculator) sumArea(shapes ...interface{}) float64 {
	var sum float64

	for _, shape := range shapes {
		switch shape.(type) {
		case circle:
			r := shape.(circle).radius
			sum += math.Pi * r * r
		case square:
			l := shape.(square).length
			sum += l * l
		}
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
