package main

import (
	"fmt"
	"os"
)

func writeBalanceFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.text", []byte(balanceText), 0644)
}

func main() {
	var accountBalance = 1000.0

	fmt.Println(".: Welcome to Go Bank :.")

	for {
		fmt.Println("How may we serve you?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Make a deposit")
		fmt.Println("3. Make a draft")
		fmt.Println("4. Exit")

		var userChoice int
		fmt.Print("Your choice: ")
		fmt.Scan(&userChoice)

		switch userChoice {
		case 1:
			fmt.Printf("\nYour balance is $%0.2f\n", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			accountBalance += depositAmount
			fmt.Printf("\nBalance updated. New amount $%0.2f\n\n", accountBalance)
			writeBalanceFile(accountBalance)
		case 3:
			fmt.Print("Your draft: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount. Overdraft not allowed.")
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Printf("\nBalance updated. New amount $%0.2f\n\n", accountBalance)
			writeBalanceFile(accountBalance)
		default:
			fmt.Println("Goodbye!")
			return
		}
	}
}
