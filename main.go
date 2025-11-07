package main

import (
	"bufio"
	"fmt"
	"os"

	"example.com/myapp/lessons"
	"example.com/myapp/lessons/bank"
	"example.com/myapp/lessons/calculators/financial/futureval"
	roi "example.com/myapp/lessons/calculators/financial/roi"
	"example.com/myapp/lessons/note"
	"example.com/myapp/lessons/todo"
)

var useDefualt bool = true

type str string

func (s str) log() {
	// Create other custom types & adding methods
	fmt.Println(s)
}

func main() {
	result := add(1, 2)
	fmt.Println("Intruducing generics: ", result)

	var name str = "Go Basics Application"

	//name = "Go Basics Application"
	name.log()

	choice := "11"

	if useDefualt {
		lessions(choice)
	} else {
		reader := bufio.NewReader(os.Stdin)

		for {
			outputText("Select a program to run:\n")
			outputText("1. Investment Calculator\n")
			outputText("2. Profit Calculator\n")
			outputText("3. Go Bank\n")
			outputText("4. Pointers\n")
			outputText("5. Structs\n")
			outputText("6. Note\n")
			outputText("7. Todo\n")
			outputText("8. Lists\n")
			outputText("9. Quiz\n")
			outputText("10 Introducing Maps\n")
			outputText("11 \"\"make\"\" function\n")
			outputText("Q. Quit\n")
			outputText("Enter choice: ")

			choice, _ = reader.ReadString('\n')
			choice = choice[:len(choice)-1] // Remove newline character

			lessions(choice)

			fmt.Println()
		}
	}
}

func lessions(choice string) {
	switch choice {
	case "1":
		futureval.New()
	case "2":
		roi.New()
	case "3":
		bank.New()
	case "4":
		lessons.Pointers() // Understanding Pointers
	case "5":
		lessons.Structs() // Structs & Custom Types
	case "6":
		note.New()
	case "7":
		todo.New()
	case "8":
		lessons.Lists()
	case "9":
		lessons.Practice()
	case "10":
		lessons.Maps()
	case "11":
		lessons.Make()
	case "Q", "q":
		outputText("Exiting the program.\n")
		return
	default:
		outputText("Invalid choice. Please try again.\n")
	}
}

func outputText(text string) {
	fmt.Printf("%v", text)
}

// func add (a, b interface {}) {
// /*
// Interfaces, Dynamic Types & Limittations
// Go accepts any kind of value here as an input.
// But as you see, I'm getting an error here now,
// that this plus operator is not defined on
// such a dynamic value type,
// because not all values support the plus operator.
// You can't for example add two structs together.
// So this is a bit too flexible.
// */
// 	return a + b;
// }

func add[T int | float64 | string](a, b T) T {
	return a + b // Gennerics
}
