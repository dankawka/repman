package repman

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "repman",
	Short: "repman is simple Git repository manager",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func RegisterCommands() {
	rootCmd.AddCommand(scanCommand)
	rootCmd.AddCommand(openCommand)
	rootCmd.AddCommand(addCommand)
	rootCmd.AddCommand(clearCommand)
}
