package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	// Revenue
	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)

	// Expenses
	fmt.Print("Enter expenses: ")
	fmt.Scan(&expenses)

	// Tax rate
	fmt.Print("Enter tax rate: ")
	fmt.Scan(&taxRate)

	// Calculate earnings before tax
	earningsBeforeTax := revenue - expenses
	fmt.Println(earningsBeforeTax)

	// Calculate earnings after tax
	earningsAfterTax := earningsBeforeTax * (1 - taxRate/100)
	fmt.Println(earningsAfterTax)

	// Calculate ratio
	ratio := earningsBeforeTax / earningsAfterTax
	fmt.Println(ratio)
}
