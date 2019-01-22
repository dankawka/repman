package prompts

import (
	"fmt"

	survey "gopkg.in/AlecAivazis/survey.v1"
)

func SingleSelect(message string, options []string) string {
	prompt := &survey.Select{
		Message:  message,
		Options:  options,
		PageSize: 25,
	}

	selectedOption := ""

	// make terminal not line wrap
	fmt.Printf("\x1b[?7l")
	survey.AskOne(prompt, &selectedOption, nil)
	// defer restoring line wrap
	defer fmt.Printf("\x1b[?7h")
	return selectedOption
}
