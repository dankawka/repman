package repman

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"

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
		FindRepositories(path)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires 1 argument - path to directory where scan should start")
		}
		return nil
	},
}
