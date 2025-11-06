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

type str string

func (s str) log() {
	// Create other custom types & adding methods
	fmt.Println(s)
}

func main() {
	var name str = "Go Basics Application"

	//name = "Go Basics Application"
	name.log()

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
		outputText("Q. Quit\n")
		outputText("Enter choice: ")

		choice, _ := reader.ReadString('\n')
		choice = choice[:len(choice)-1] // Remove newline character

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
		case "Q", "q":
			outputText("Exiting the program.\n")
			return
		default:
			outputText("Invalid choice. Please try again.\n")
		}

		fmt.Println()
	}
}

func outputText(text string) {
	fmt.Printf("%v", text)
}
