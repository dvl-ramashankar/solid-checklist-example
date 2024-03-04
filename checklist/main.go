package main

import (
	"fmt"
	"log"
	"time"
)

const (
	Client = "one"
	Test   = "test"
)

func main() {
	var add = Addition{value1: 5, value2: 5}
	fmt.Println("Addition: ", add.addition())
	var flag1 = checkFlag1()
	var flag2 = checkFlag2()
	if flag1 && flag2 {
		subValue, err := subtract(5, 8)
		if err != nil {
			log.Println(err.Error())
		}
		fmt.Println("Subtraction: ", subValue)
	}
	fmt.Println("Switch Block : ", handleSwitchBlock(Client))
	fmt.Println(Test)
	month, err := FetchMonthName("22-02-2024")
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("Time :", month)
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

func subtract(a, b int) (int, error) {
	if a == 0 && b == 0 {
		return 0.0, fmt.Errorf("invalid arguments value1 : %d and value2: %d", a, b)
	}
	return a - b, nil
}

func checkFlag1() bool {
	return true
}

func checkFlag2() bool {
	return true
}

func handleSwitchBlock(value string) int {
	switch value {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	default:
		return 0
	}
}

func FetchMonthName(dt string) (string, error) {
	// Parse the date string
	date, err := time.Parse("02-01-2006", dt)
	if err != nil {
		return "", fmt.Errorf("given date is %s. please provide date in format of DD-MM-YYYY", dt)
	}

	// Get the month name
	monthName := date.Month().String()

	return monthName, nil
}
