// You can write Go code here!
package investmentcalculator

import (
	"fmt"
	"math"
)

func main() {
	// Application entry point

	//fmt.Print("Hello World!")
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	fmt.Print("Enter the investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter the expected return rate (in %): ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter the number of years: ")
	fmt.Scan(&years)

	// Calculate future value
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Printf("Future value of the investment: %.2f\n", futureValue)
	fmt.Printf("Future real value of the investment (adjusted for inflation): %.2f\n", futureRealValue)
	// To verify whether this is correct or not, opening up some future value calculator online: https://bitl.to/5HQ2

	// Example output:
}

// go run app.go
// Hello World!
// go build app.go
// ./app
// Hello World!

// go mod init
