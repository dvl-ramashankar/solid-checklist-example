package main

import "fmt"

const Client = "Demo"

func main() {
	fmt.Println(Client)
	add := Addition{value1: 5, value2: 5}
	fmt.Println(add.addition())
}

type Addition struct {
	value1 float64
	value2 float64
}

func (a Addition) addition() float64 {
	return sum(a.value1, a.value2)
}

type Calculation interface {
	sum(a, b float64) float64
}

func sum(a, b float64) float64 {
	return a + b
}
