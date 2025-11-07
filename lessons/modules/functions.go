package modules

type transformFn func(int) int

// type anotherFn func(int, []string, map[string][]int) ([]int, string)
var Numbers = []int{1, 2, 3, 4}

func Functions() {
	// Module intruduction

	// fmt.Println("Functions as values & function types.")

	// doubled := transformNumbers(&Numbers, double)
	// tripled := transformNumbers(&Numbers, triple)

	// fmt.Println(doubled)
	// fmt.Println(tripled)

	// Return()

	// Anonymous()

	//Recursion()

	Variadic()
}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumber := []int{}

	for _, val := range *numbers {
		dNumber = append(dNumber, transform(val))
	}

	return dNumber
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
