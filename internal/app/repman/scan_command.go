package repman

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/dankawka/repman/internal/pkg/repofinder"
	settingsmanager "github.com/dankawka/repman/internal/pkg/settings_manager"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

func ask(repositories []repofinder.Repository) []repofinder.Repository {
	chosenOptions := []string{}

	repositoriesNormalized := []string{}

	for _, repository := range repositories {
		repositoriesNormalized = append(repositoriesNormalized, fmt.Sprintf("Path: %s | Origin: %s", repository.Path, repository.Origin))
	}

	prompt := &survey.MultiSelect{
		Message:  "Which repository you would like to save:",
		Options:  repositoriesNormalized,
		Default:  repositoriesNormalized,
		PageSize: 25,
	}

	survey.AskOne(prompt, &chosenOptions, nil)

	if len(chosenOptions) == 0 {
		color.Yellow("You chose 0 options. Exiting application.")
	}

	chosenRepositories := []repofinder.Repository{}
	for _, repository := range repositories {
		for _, option := range chosenOptions {
			if strings.Contains(option, repository.Origin) && strings.Contains(option, repository.Path) {
				chosenRepositories = append(chosenRepositories, repository)
			}
		}
	}

	return chosenRepositories
}

var scanCommand = &cobra.Command{
	Use:   "scan [PATH]",
	Short: "Scans recursively folder for Git repositories",
	Run: func(cmd *cobra.Command, args []string) {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		path := path.Join(dir, args[0])
		repositories := repofinder.FindRepositories(path)
		chosenRepositories := ask(repositories)
		settingsmanager.SaveRepositories(chosenRepositories)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires 1 argument - path to directory where scan should start")
		}
		return nil
	},
}
