package main

import "fmt"

type shape interface {
	area() float64
	volume() float64
}

func areaSum(shapes ...shape) float64 {
	var sum float64
	for _, s := range shapes {
		sum += s.area()
	}
	return sum
}

func areaVolumeSum(shapes ...shape) float64 {
	var sum float64
	for _, s := range shapes {
		sum += s.area() + s.volume()
	}
	return sum
}

type square struct {
	length float64
}

func (s square) area() float64 {
	return s.length * s.length
}

func (s square) volume() float64 {
	return 0
}

type cube struct {
	length float64
}

func (c cube) area() float64 {
	return 6.0 * c.length * c.length
}

func (c cube) volume() float64 {
	return c.length * c.length * c.length
}

func main() {
	s := square{length: 1.0}
	c := cube{length: 2.0}
	fmt.Println(areaSum(s, c))
	fmt.Println(areaVolumeSum(s, c))
}
