package bank

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	// Placeholder for file reading logic
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 1000, errors.New("Failed to read balance file.")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("Failed to parse stored balance value.")
	}

	return balance, nil
}

func writeBalanceToFile(balance float64) {
	// Placeholder for file writing logic
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func Run() {
	var accountBalance, err = getBalanceFromFile()
	// Placeholder for bank application logic
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("----------- ")
		//return
		//panic("Can't continue, sorry!")
	}

	for {
		fmt.Println("Welcome to Go Bank!")
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		//wantsCheckBalance := choice == 1
		switch choice {
		case 1:
			fmt.Printf("Your balance is: $%.2f\n", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				//return
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			fmt.Print("Withdrawal amount: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				return
			}

			if withdrawalAmount > accountBalance {
				fmt.Print("Invalid amount. You can't withdraw more than you have.")
				return
			}

			accountBalance -= withdrawalAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thank you for using Go Bank!")
			return
			//break
		}
	}
}
