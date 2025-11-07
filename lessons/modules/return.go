package modules

import (
	"fmt"
)

func Return() {
	//numbers := []int{1, 2, 3, 4}
	// doubled := transformNumbers(&numbers, double)
	// tripled := transformNumbers(&numbers, triple)

	// fmt.Println(doubled)
	// fmt.Println(tripled)
	fmt.Println("Returning functions as values.")
	moreNumbers := []int{5, 1, 2}

	transformerFn1 := getTransformerFunction(&Numbers)
	transformerFn2 := getTransformerFunction(&moreNumbers)

	transformeredNumbers := transformNumbers(&Numbers, transformerFn1)
	moreTransformedNumbers := transformNumbers(&moreNumbers, transformerFn2)

	fmt.Println(transformeredNumbers)
	fmt.Println(moreTransformedNumbers)
}

func getTransformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}
