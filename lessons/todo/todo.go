package todo

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/myapp/lessons/note/noted"
	"example.com/myapp/lessons/todo/todod"
)

type saver interface {
	Save() error
}

func New() {
	// This is a placeholder file to demonstrate the note lesson.
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todod.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := noted.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	todo.Display()
	err = saveData(todo)

	if err != nil {
		return
	}

	userNote.Display()
	err = saveData(userNote)

	if err != nil {
		return
	}
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return err
	}

	fmt.Println("Saving the note succeeded!")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")
	return title, content
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
