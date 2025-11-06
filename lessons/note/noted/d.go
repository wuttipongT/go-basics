package noted

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	title    string
	content  string
	createAt time.Time
}

func (n Note) Display() {
	fmt.Printf("Your note titled %v has the following content:\n\n%v", n.title, n.content)
}

func (n Note) Save() error {
	fileName := strings.ReplaceAll(n.title, " ", "_")
	fileName = strings.ToLower(fileName)

	json, err := json.Marshal(n)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("Invalid input.")
	}
	return Note{
		title:    title,
		content:  content,
		createAt: time.Now(),
	}, nil
}
