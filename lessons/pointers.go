package lessons

import "fmt"

func Pointers() {
	age := 35 // regular variable

	var agePointer *int // pointer variable declaration

	agePointer = &age // pointer variable initialization

	fmt.Println("Age:", *agePointer)

	//adultYear := getAdultYear(agePointer)
	editAgeToAdultYears(agePointer)
	fmt.Println(age)
	//* , an asterisk if want to get value from pointer
	//& , an ampersand
}

func editAgeToAdultYears(age *int) {
	// return *age - 18
	*age = *age - 18
}
