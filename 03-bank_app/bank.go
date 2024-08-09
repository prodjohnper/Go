package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

// Function to get balance from file
func getBalanceFromFile() (float64, error) {
	// Read data from file
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 0, errors.New("Account balance couldn't be retrieved, try again later.")
	}

	// Convert data to string
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 0, errors.New("An error has ocurred while parsing data, try again later")
	}

	// Return balance for latter use
	return balance, nil
}

// Function to save and update accountBalance in a file,
func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	// Balance
	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		// Print divider
		fmt.Println(`- - - - - - - - - - - - - - - -`)
	}

	// Output greeting message
	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("Choose an option from the menu")

		// Print divider
		fmt.Println(`- - - - - - - - - - - - - - - -`)

		// Output options menu
		greeting()

		// Print divider
		fmt.Println(`- - - - - - - - - - - - - - - -`)

		// Get user input
		var userChoice int
		fmt.Print("Your choice: ")
		fmt.Scan(&userChoice)

		// Print divider
		fmt.Println(`- - - - - - - - - - - - - - - -`)

		if userChoice == 1 {
			fmt.Printf("Your balance is: %.2f", accountBalance)
		} else if userChoice == 2 {
			// Deposit amount
			var depositAmount float64

			// Get user input
			fmt.Print("Amount to deposit: ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				// Check if deposit amount is 0 or a negative num
				fmt.Println("Invalid amount, try again.")
				continue
			}

			accountBalance += depositAmount
			writeBalanceToFile(accountBalance)

			fmt.Printf("Your new balance is: %.2f", accountBalance)

		} else if userChoice == 3 {
			// Withdraw amount
			var withdrawAmount float64

			fmt.Print("Amount to withdraw: ")
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				// Print divider
				fmt.Println(`- - - - - - - - - - - - - - - -`)

				fmt.Println("Invalid amount, try again.")
				continue
			}

			if withdrawAmount > accountBalance {
				// Print divider
				fmt.Println(`- - - - - - - - - - - - - - - -`)

				fmt.Println("Not enough funds. Your current balance is:", accountBalance)
				continue
			}

			accountBalance -= withdrawAmount
			writeBalanceToFile(accountBalance)

			// Print divider
			fmt.Println(`- - - - - - - - - - - - - - - -`)

			fmt.Printf("Your new balance is: %.2f", accountBalance)
		} else if userChoice == 4 {
			fmt.Println("Thanks for using our services! See you soon!")
			break
		} else {
			fmt.Print("Invalid selection, try again.")
		}
		// Print divider
		fmt.Printf("\n")
		fmt.Println(`- - - - - - - - - - - - - - - -`)
	}
}
