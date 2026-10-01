package main

import (
	"fmt"
	"math"
)

func main() {
	const inflation float64 = 2.5
	var rate float64
	var years float64
	var amount float64

fmt.Print("Enter the initial amount: ")
	fmt.Scanln(&amount)

	fmt.Print("Enter the number of years: ")
	fmt.Scanln(&years)

	rate = inflation / 100

	futureValue := amount * math.Pow(1+rate, years)

	fmt.Printf("The future value after %.0f years is: %.2f\n", years, futureValue)

}