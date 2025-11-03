// You can write Go code here!
package main

import (
	"fmt"
	"math"
)

func main() {
	// Application entry point

	//fmt.Print("Hello World!")

	var investmentAmount float64 = 10000
	years := 10.0
	expectedReturnRate := 5.5

	// Calculate future value
	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	fmt.Printf("Future value of the investment: %.2f\n", futureValue)
	// To verify whether this is correct or not, opening up some future value calculator online: https://bitl.to/5HQ2
}

// go run app.go
// Hello World!
// go build app.go
// ./app
// Hello World!

// go mod init
