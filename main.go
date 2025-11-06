package main

import (
	"bufio"
	"fmt"
	"os"

	"example.com/myapp/lessons"
	"example.com/myapp/lessons/bank"
	"example.com/myapp/lessons/calculators"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		outputText("Select a program to run:\n")
		outputText("1. Investment Calculator\n")
		outputText("2. Profit Calculator\n")
		outputText("3. Go Bank\n")
		outputText("4. Pointers\n")
		outputText("5. Structs\n")
		outputText("Q. Quit\n")
		outputText("Enter choice: ")

		choice, _ := reader.ReadString('\n')
		choice = choice[:len(choice)-1] // Remove newline character

		switch choice {
		case "1":
			calculators.Investment()
		case "2":
			calculators.Profit()
		case "3":
			bank.Run()
		case "4":
			lessons.Pointers() // Understanding Pointers
		case "5":
			lessons.Structs() // Structs & Custom Types
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
