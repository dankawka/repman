package repman

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/dankawka/repman/internal/pkg/repofinder"
	settingsmanager "github.com/dankawka/repman/internal/pkg/settings_manager"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

func askToOpen(repositories []repofinder.Repository) {
	repositoriesNormalized := []string{}
	for _, repository := range repositories {
		repositoriesNormalized = append(repositoriesNormalized, fmt.Sprintf("Path: %s | Origin: %s", repository.Path, repository.Origin))
	}

	prompt := &survey.Select{
		Message:  "Which repository you would like to open:",
		Options:  repositoriesNormalized,
		PageSize: 25,
	}

	var selectedOption = ""

	// make terminal not line wrap
	fmt.Printf("\x1b[?7l")
	survey.AskOne(prompt, &selectedOption, nil)
	// defer restoring line wrap
	defer fmt.Printf("\x1b[?7h")

	if len(selectedOption) == 0 {
		color.Yellow("No option chosen, exiting app.")
		os.Exit(0)
	}

	chosenRepository := repofinder.Repository{}
	for _, repository := range repositories {
		if strings.Contains(selectedOption, repository.Origin) && strings.Contains(selectedOption, repository.Path) {
			chosenRepository = repository
		}
	}

	color.Green("Opening %s", chosenRepository.Path)
	command := exec.Command("code", chosenRepository.Path)
	command.Start()
	os.Exit(0)
}

var openCommand = &cobra.Command{
	Use:   "open",
	Short: "Shows list of saved repositories and opens selected",
	Run: func(cmd *cobra.Command, args []string) {
		repos, err := settingsmanager.GetListOfRepositories()
		if err != nil {
			os.Exit(1)
		}
		askToOpen(repos)
	},
}
