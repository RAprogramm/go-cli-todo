package todo

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

func Input(questtion string) string {
	prompt := promptui.Prompt{
		Label: questtion,
	}

	answer, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return answer
}

func Select() string {
	prompt := promptui.Select{
		Label: "What do you want?",
		Items: []string{"Create a new task", "Delete a task", "Mark task as completed", "Exit"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return result
}
