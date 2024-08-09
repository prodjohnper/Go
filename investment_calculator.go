package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 2.5 // inflation rate

	var investmentAmount float64 = 1000 // initial investment amount
	expectedReturnRate := 5.5           // expected return rate
	years := 10.0                       // number of years

	// future value of the investment
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// print the future value & future real value
	fmt.Println("Future value:", futureValue)
	fmt.Println("Future real value:", futureRealValue)
}
