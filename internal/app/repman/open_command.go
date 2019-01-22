package repman

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/dankawka/repman/internal/pkg/models"
	"github.com/dankawka/repman/internal/pkg/prompts"
	settingsmanager "github.com/dankawka/repman/internal/pkg/settings_manager"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func askToOpen(repositories []models.Repository) {
	repositoriesNormalized := []string{}
	for _, repository := range repositories {
		repositoriesNormalized = append(repositoriesNormalized, fmt.Sprintf("Path: %s | Origin: %s", repository.Path, repository.Origin))
	}

	selectedOption := prompts.SingleSelect("Which repository you would like to open:", repositoriesNormalized)

	if len(selectedOption) == 0 {
		color.Yellow("No option chosen, exiting app.")
		os.Exit(0)
	}

	chosenRepository := models.Repository{}
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

		if len(repos) == 0 {
			color.Yellow("No repositories saved, use option 'add' or 'scan'")
			os.Exit(0)
		}

		askToOpen(repos)
	},
}
