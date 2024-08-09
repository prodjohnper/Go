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

	earningsAfterTax := earningsBeforeTax * (1 - taxRate/100)
	fmt.Println(earningsAfterTax)

	// Ratio
	ratio := earningsBeforeTax / earningsAfterTax
	fmt.Println(ratio)
}
