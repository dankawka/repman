package repman

import (
	"github.com/dankawka/repman/internal/pkg/models"
	settingsmanager "github.com/dankawka/repman/internal/pkg/settings_manager"
	"github.com/spf13/cobra"
)

var clearCommand = &cobra.Command{
	Use:   "clear",
	Short: "Clears saved list of repositories",
	Run: func(cmd *cobra.Command, args []string) {
		emptyList := []models.Repository{}
		settingsmanager.SaveRepositories(emptyList)
	},
}
