package profitcalculator

import (
	"fmt"
)

func Run() {
	var revenue float64
	var expenses float64
	var txtRate float64

	revenue = getUserInput("Revenue: ")
	expenses = getUserInput("Expenses: ")
	txtRate = getUserInput("Tax Rate (%): ")

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
}

func calculateFinancials(revenue, expenses, txtRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - txtRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}
