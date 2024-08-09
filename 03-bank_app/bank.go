package main

import (
	"fmt"
	"os"
)

// Function to save and update accountBalance in a file,
func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}

func main() {
	// Balance
	var accountBalance float64 = 1000.00

	// Output greeting message
	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("Choose an option from the menu")

		// Print divider
		fmt.Println(`- - - - - - - - - - - - - - - -`)

		// Output options menu
		fmt.Println("1. Balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Exit")

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
			fmt.Println("Invalid selection, try again.")
		}
		// Print divider
		fmt.Printf("\n")
		fmt.Println(`- - - - - - - - - - - - - - - -`)
	}
}
