package repman

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/dankawka/repman/internal/pkg/repofinder"
	settingsmanager "github.com/dankawka/repman/internal/pkg/settings_manager"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

func askToSave(repositories []repofinder.Repository) []repofinder.Repository {
	chosenOptions := []string{}
	chosenRepositories := []repofinder.Repository{}
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
		return chosenRepositories
	}

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
		path, err := filepath.Abs(args[0])
		if err != nil {
			color.Red(err.Error())
			os.Exit(1)
		}
		repositories, err := repofinder.FindRepositories(path)
		if err != nil {
			color.Red(err.Error())
			os.Exit(1)
		}
		chosenRepositories := askToSave(repositories)

		if len(chosenRepositories) == 0 {
			color.Yellow("You chose 0 options. Exiting application.")
			os.Exit(0)
		}

		currentRepositories, _ := settingsmanager.GetListOfRepositories()

		extended := currentRepositories

		for _, r := range chosenRepositories {
			alreadyExists := false
			for _, cr := range currentRepositories {
				if r.Origin == cr.Origin && r.Path == cr.Path {
					alreadyExists = true
				}
			}
			if alreadyExists {
				alreadyExists = false
			} else {
				extended = append(extended, r)
			}
		}

		settingsmanager.SaveRepositories(extended)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires 1 argument - path to directory where scan should start")
		}
		return nil
	},
}
