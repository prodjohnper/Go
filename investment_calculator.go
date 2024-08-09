package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64 = 1000 // initial investment amount
	expectedReturnRate := 5.5           // expected return rate
	years := 10.0                       // number of years

	// future value of the investment
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	// print the future value
	fmt.Println("Future value:", futureValue)
}
