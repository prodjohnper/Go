package main

import (
	"fmt"
)

func main() {
	revenue := getUserInput("Revenue: ")   // Revenue
	expenses := getUserInput("Expenses: ") // Expenses
	taxRate := getUserInput("Tax rate: ")  // Tax rate

	// Print divider
	fmt.Println(`- - - - - - - - - - - - - - - -`)

	// Calculate earnings before/after tax and ratio
	earningsBeforeTax, earningsAfterTax, ratio :=
		calculateFinancials(revenue, expenses, taxRate)

	// EBT, profit and ratio calculation results
	fmt.Printf("EBT: %.2f\n", earningsBeforeTax)
	fmt.Printf("Profit: %.2f\n", earningsAfterTax)
	fmt.Printf("Ratio: %.2f\n", ratio)
}

// Function that gets user input and sets pointer to respective variable
func getUserInput(text string) float64 {
	var userInput float64

	fmt.Print(text)
	fmt.Scan(&userInput)

	return userInput
}

// Function that calculates EBT, profit and ratio
func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earningsBeforeTax := revenue - expenses
	earningsAfterTax := earningsBeforeTax * (1 - taxRate/100)
	ratio := earningsBeforeTax / earningsAfterTax

	return earningsBeforeTax, earningsAfterTax, ratio
}
