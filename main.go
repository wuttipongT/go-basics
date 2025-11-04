package main

import (
	"bufio"
	"fmt"
	"os"

	bank "example.com/myapp/essentials/bank"
	investmentcalculator "example.com/myapp/essentials/investment_calculator"
	profitcalculator "example.com/myapp/essentials/profit_calculator"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		outputText("Select a program to run:\n")
		outputText("1. Investment Calculator\n")
		outputText("2. Profit Calculator\n")
		outputText("3. Go Bank\n")
		outputText("Q. Quit\n")
		outputText("Enter choice: ")

		choice, _ := reader.ReadString('\n')
		choice = choice[:len(choice)-1] // Remove newline character

		switch choice {
		case "1":
			investmentcalculator.Run()
		case "2":
			profitcalculator.Run()
		case "3":
			bank.Run()
		case "Q", "q":
			outputText("Exiting the program.\n")
			return
		default:
			outputText("Invalid choice. Please try again.\n")
		}

		fmt.Println()
	}
}

func outputText(text string) {
	fmt.Printf("%v", text)
}
