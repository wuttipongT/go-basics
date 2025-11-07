package modules

import "fmt"

func Variadic() {
	fmt.Println("Using variadic functions.")

	sum := sumup(1, 10, 15, 40, -5)
	anotherSum := sumup(1, Numbers...)

	fmt.Println(sum)
	fmt.Println(anotherSum)
}

func sumup(startingValue int, numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val // sum = sum + val
	}

	return sum
}
