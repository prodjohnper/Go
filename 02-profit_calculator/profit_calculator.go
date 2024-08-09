package main

import (
	"errors"
	"fmt"
	"os"
)

// File name
const resultsFile string = "results.txt"

// Function that gets user input and sets pointer to respective variable
func getUserInput(text string) (float64, error) {
	var userInput float64

	fmt.Print(text)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("Invalid amount, try again!")
	}

	return userInput, nil
}

// Function that calculates EBT, profit and ratio
func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earningsBeforeTax := revenue - expenses
	earningsAfterTax := earningsBeforeTax * (1 - taxRate/100)
	ratio := earningsBeforeTax / earningsAfterTax

	return earningsBeforeTax, earningsAfterTax, ratio
}

// Function to save and update calculations results into file
func writeBalanceToFile(earningsBeforeTax, earningsAfterTax, ratio float64) {
	data := fmt.Sprintf(
		"Ebt: %.2f\nProfit: %.2f\nRatio: %.2f\n", earningsBeforeTax, earningsAfterTax, ratio)

	os.WriteFile(resultsFile, []byte(data), 0644)
}

// Function to handle errors
func validateData(err error) bool {
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func main() {
	println("Welcome!")
	// Print divider
	fmt.Println(`- - - - - - - - - - - - - - - -`)

	for {
		revenue, err := getUserInput("Revenue: ") // Revenue
		if !validateData(err) {
			continue
		}

		expenses, err := getUserInput("Expenses: ") // Expenses
		if !validateData(err) {
			continue
		}

		taxRate, err := getUserInput("Tax rate: ") // Tax rate
		if !validateData(err) {
			continue
		}

		// Print divider
		fmt.Println(`- - - - - - - - - - - - - - - -`)

		// Calculate earnings before/after tax and ratio
		earningsBeforeTax, earningsAfterTax, ratio :=
			calculateFinancials(revenue, expenses, taxRate)

		writeBalanceToFile(earningsBeforeTax, earningsAfterTax, ratio)

		// EBT, profit and ratio calculation results
		fmt.Printf("EBT: %.2f\n", earningsBeforeTax)
		fmt.Printf("Profit: %.2f\n", earningsAfterTax)
		fmt.Printf("Ratio: %.2f\n", ratio)

		// Print divider
		fmt.Println(`- - - - - - - - - - - - - - - -`)
	}
}
