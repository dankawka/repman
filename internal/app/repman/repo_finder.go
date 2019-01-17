package repman

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"

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

	return string(out)
}

type Repository struct {
	path   string
	origin string
}

func FindRepositories(path string) {
	repositories := []Repository{}
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		if !info.IsDir() {
			// not a directory, skiping
			return nil
		}

		found := checkIfGitDirectoryExists(path)

		if found {
			remote := getOrigin(path)
			repositories = append(repositories, Repository{origin: remote, path: path})
		}

		return nil
	})

	color.Yellow("Found %v repositories.", len(repositories))

}
