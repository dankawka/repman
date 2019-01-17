package repofinder

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

func checkIfGitDirectoryExists(dirPath string) bool {
	gitPath := path.Join(dirPath, ".git")
	_, err := os.Stat(gitPath)
	if err != nil {
		// .git folder does not exist
		return false
	}

	return true
}

func getOrigin(path string) string {
	command := exec.Command("git", "remote", "get-url", "origin")
	command.Dir = path

	out, err := command.Output()

	if err != nil {
		fmt.Println("Something went wrong :(")
		return ""
	}

	return strings.TrimSpace(string(out))
}

type Repository struct {
	Path   string
	Origin string
}

func FindRepositories(path string) []Repository {
	repositories := []Repository{}
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			color.Red("Could not access path %s", path)
			return err
		}

		if !info.IsDir() {
			// not a directory, skiping
			return nil
		}

		found := checkIfGitDirectoryExists(path)

		if found {
			remote := getOrigin(path)
			repositories = append(repositories, Repository{Origin: remote, Path: path})
		}

		return nil
	})

	if len(repositories) == 0 {
		color.Yellow("Could not find repositoties under provided path.")
		return repositories
	}

	color.Green("Found %v repositories.", len(repositories))

	return repositories
}
