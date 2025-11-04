// You can write Go code here!
package investmentcalculator

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	// Application entry point

	//fmt.Print("Hello World!")
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	fmt.Print("Enter the investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter the expected return rate (in %): ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter the number of years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	// Calculate future value
	//futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	//futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Printf("Future value of the investment: %.2f\n", futureValue)
	fmt.Printf("Future real value of the investment (adjusted for inflation): %.2f\n", futureRealValue)
	// To verify whether this is correct or not, opening up some future value calculator online: https://bitl.to/5HQ2

	// Example output:
}

func outputText(text string) {
	fmt.Println(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	// Placeholder for future value calculations
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)

	return fv, rfv
}

// go run app.go
// Hello World!
// go build app.go
// ./app
// Hello World!

// go mod init
