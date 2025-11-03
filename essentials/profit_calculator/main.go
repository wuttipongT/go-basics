package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var txtRate float64

	fmt.Print("Enter the total revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter the total expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter the tax rate (in %): ")
	fmt.Scan(&txtRate)

	ebt := revenue - expenses
	profit := ebt * (1 - txtRate/100)
	ratio := ebt / profit

	fmt.Printf("Profit before tax: %d\n", ebt)
	fmt.Printf("Tax amount: %.2f\n", profit)
	fmt.Printf("Profit after tax: %d\n", ratio)
}
