package note

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/myapp/lessons/note/noted"
)

func New() {
	// This is a placeholder file to demonstrate the note lesson.
	title, content, err := getNoteData()

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := noted.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return
	}

	fmt.Println("Saving the note succeeded!")
}

func getNoteData() (string, string, error) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")
	return title, content, nil
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
