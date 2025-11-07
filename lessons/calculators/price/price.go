package price

import (
	"fmt"

	"example.com/jeff/go-basics/lessons/calculators/price/filemanager"
	"example.com/jeff/go-basics/lessons/calculators/price/prices"
)

func New() {
	fmt.Println("Price Calculator.")
	// prices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.7, 0.1, 0.15}

	//var result map[float64][]float64 = make(map[float64][]float64)
	//result := make(map[float64][]float64)

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		//cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		err := priceJob.Process()

		if err != nil {
			fmt.Println("Cloud not process job")
			fmt.Println(err)
		}
	}

	//fmt.Println(result)
}
