package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64   // declare investment amount var
	var expectedReturnRate float64 // expected return rate
	var years float64              // number of years

	// inflation rate
	const inflationRate float64 = 2.5

	// Investment amount
	fmt.Print("Enter investment amount: ") // prompt user to enter investment amount
	fmt.Scan(&investmentAmount)            // read the investment amount

	// Expected return rate
	fmt.Print("Enter expected return rate: ") // prompt user to enter expected return rate
	fmt.Scan(&expectedReturnRate)             // read the expected return rate

	// Number of years
	fmt.Print("Enter number of years: ") // prompt user to enter number of years
	fmt.Scan(&years)                     // read the number of years

	// future value of the investment
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// Print divider
	fmt.Println(`- - - - - - - - - - - - - - - -`)

	// print the future value & future real value
	fmt.Printf("Future value: $%.2f\n", futureValue)
	fmt.Printf("Future real value: $%.2f\n", futureRealValue)
}
