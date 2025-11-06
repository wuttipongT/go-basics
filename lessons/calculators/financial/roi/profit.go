package roi

import (
	"errors"
	"fmt"
	"os"
)

// Goals
// 1) Validate user input
//  => Show error message & exit if invalid input is provided
//  - No negative numbers
//  - Not 0
// 2) Store calculatted results into file

func New() {
	// var revenue float64
	// var expenses float64
	// var txtRate float64

	revenue, err := getUserInput("Revenue: ") // 1000

	if err != nil {
		fmt.Println(err)
		return
	}

	expenses, err := getUserInput("Expenses: ") // 50

	if err != nil {
		fmt.Println(err)
		return
	}

	txtRate, err := getUserInput("Tax Rate (%): ") // 25

	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Print("Enter the total revenue: ")
	// fmt.Scan(&revenue)
	// fmt.Print("Enter the total expenses: ")
	// fmt.Scan(&expenses)
	// fmt.Print("Enter the tax rate (in %): ")
	// fmt.Scan(&txtRate)

	// ebt := revenue - expenses
	// profit := ebt * (1 - txtRate/100)
	// ratio := ebt / profit

	ebt, profit, ratio := calculateFinancials(revenue, expenses, txtRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)

	storeResults(ebt, profit, ratio)
}

func storeResults(ebt, profit, ratio float64) {
	// Placeholder for file writing logic
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}

func calculateFinancials(revenue, expenses, txtRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - txtRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("Value must be a positive number.")
	}

	return userInput, nil
}
