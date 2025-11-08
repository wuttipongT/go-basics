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
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))
	//var result map[float64][]float64 = make(map[float64][]float64)
	//result := make(map[float64][]float64)

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		//cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[index], errorChans[index])

		// if err != nil {
		// 	fmt.Println("Cloud not process job")
		// 	fmt.Println(err)
		// }
	}

	for index := range taxRates {
		select {
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[index]:
			fmt.Println("Done!")
		}
	}

	// for _, errorChan := range errorChans {
	// 	<-errorChan
	// }

	// for _, doneChan := range doneChans {
	// 	<-doneChan
	// }
	//fmt.Println(result)
}
