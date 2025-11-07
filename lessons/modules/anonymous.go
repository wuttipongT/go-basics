package modules

import "fmt"

func Anonymous() {
	fmt.Println("Intrducing anonymous functions.")

	double := createTransformer(2)
	triple := createTransformer(3)

	transformed := transformNumbers(&Numbers, func(number int) int {
		return number * 2
	})

	doubled := transformNumbers(&Numbers, double)
	tripled := transformNumbers(&Numbers, triple)

	fmt.Println(transformed)
	fmt.Println(doubled)
	fmt.Println(tripled)
}

func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
