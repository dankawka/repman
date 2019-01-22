package prompts

import (
	"fmt"

	survey "gopkg.in/AlecAivazis/survey.v1"
)

func MultiselectPrompt(message string, options, defaultOptions []string) []string {
	prompt := &survey.MultiSelect{
		Message:  message,
		Options:  options,
		Default:  defaultOptions,
		PageSize: 10,
	}

	chosenOptions := []string{}
	// make terminal not line wrap
	fmt.Printf("\x1b[?7l")
	survey.AskOne(prompt, &chosenOptions, nil)
	// defer restoring line wrap
	defer fmt.Printf("\x1b[?7h")
	return chosenOptions
}
