package lessons

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthday  string
	createdAt time.Time
}

func Structs() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthday := getUserData("Please enter your birthdate (MM-DD-YYYY): ")

	var appUser user

	appUser = user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthday:  userBirthday,
		createdAt: time.Now(),
	}

	// ... do something awesome with that gathrered data!

	outputUserDetails(&appUser) // Struct & Pointers
}

func outputUserDetails(u *user) {
	// ... (*u).firstName
	fmt.Println(u.firstName, u.lastName, u.birthday)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
