// You can write Go code here!
package investmentcalculator

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func Run() {
	// Application entry point

	//fmt.Print("Hello World!")
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	outputText("Starting Amount (PV): ") // 10000
	fmt.Scan(&investmentAmount)

	outputText("Interest Rate (I/Y): ")
	fmt.Scan(&expectedReturnRate)

	outputText("Number of Periods (N): ") // 10
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	// Calculate future value
	//futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	//futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Printf("Future Value: %.2f\n", futureValue)
	fmt.Printf("Future Real Value: %.2f\n", futureRealValue)
	// To verify whether this is correct or not, opening up some future value calculator online: https://bitl.to/5HQ2

	// Example output:
}

func outputText(text string) {
	fmt.Printf("%v", text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	// Placeholder for future value calculations
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)

	//return fv, rfv
	return
}

// go run app.go
// Hello World!
// go build app.go
// ./app
// Hello World!

// go mod init
