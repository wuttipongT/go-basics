package price

import (
	"fmt"

	"example.com/jeff/go-basics/lessons/calculators/price/prices"
)

func New() {
	fmt.Println("Price Calculator.")
	// prices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.7, 0.1, 0.15}

	//var result map[float64][]float64 = make(map[float64][]float64)
	//result := make(map[float64][]float64)

	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJob(taxRate)
		priceJob.Process()
	}

	//fmt.Println(result)
}
