package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	var floats []float64

	for _, stringVal := range strings {
		floatVal, err := strconv.ParseFloat(stringVal, 64)

		if err != nil {
			//fmt.Println("Converting prie to float failed.")
			//fmt.Println(err)
			//file.Close()
			return nil, errors.New("Failed to convert string to float.")
		}

		floats = append(floats, floatVal)
		//prices[stringIndex] = floatPrice
	}

	return floats, nil
}
