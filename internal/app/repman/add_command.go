package repman

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/dankawka/repman/internal/pkg/repofinder"
	settingsmanager "github.com/dankawka/repman/internal/pkg/settings_manager"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var addCommand = &cobra.Command{
	Use:   "add [PATH]",
	Short: "Adds repository under path to list of repositories",
	Run: func(cmd *cobra.Command, args []string) {
		path, err := filepath.Abs(args[0])
		if err != nil {
			color.Red(err.Error())
			os.Exit(1)
		}

		repository, err := repofinder.GetRepositoryFromPath(path)
		if err != nil {
			color.Red(err.Error())
			os.Exit(1)
		}

		exists := settingsmanager.CheckIfAlreadySaved(repository)

		if exists {
			color.Yellow("Repository %s is already saved, nothing to do, exiting...", repository.Origin)
			os.Exit(0)
		}

		color.Green("Found repository %s, adding to the list", repository.Origin)

		settingsmanager.AppendRepository(repository)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires 1 argument - path to directory with Git repository to save")
		}
		return nil
	},
}
