package lessons

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func Make() {
	userNames := make([]string, 2, 5)
	//userNames := []string{} // It'll get an error for userName[0]...

	userNames[0] = "Julie"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)

	courseRatings := make(floatMap, 3)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.7

	courseRatings.output()

	//fmt.Println(courseRatings)

	for index, value := range userNames {
		// ...
		fmt.Println("Index:", index)
		fmt.Println("Value:", value)
	}

	for index, value := range courseRatings {
		// ...
		fmt.Println("Key:", index)
		fmt.Println("Value:", value)
	}
}
