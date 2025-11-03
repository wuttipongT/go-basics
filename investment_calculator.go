// You can write Go code here!
package main

import "fmt"

func main() {
	// Application entry point

	//fmt.Print("Hello World!")

	var investmentAmount = 1000
	var expectedReturnRate = 5.5
	var years = 10

	// Calculate future value
	var futureValue = investmentAmount * (1 + expectedReturnRate/100) * years
	fmt.Printf("Future value of the investment: %.2f\n", futureValue)
}

// go run app.go
// Hello World!
// go build app.go
// ./app
// Hello World!

// go mod init
