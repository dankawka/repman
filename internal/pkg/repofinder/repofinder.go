package repofinder

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/dankawka/repman/internal/pkg/models"
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

func FindRepositories(path string) ([]models.Repository, error) {
	repositories := []models.Repository{}
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("Could not access path %s", path)
		}

		if !info.IsDir() {
			// not a directory, skiping
			return nil
		}

		found := checkIfGitDirectoryExists(path)

		if found {
			remote := getOrigin(path)
			repositories = append(repositories, models.Repository{Origin: remote, Path: path})
		}

		return nil
	})

	if err != nil {
		return repositories, err
	}

	if len(repositories) == 0 {
		return repositories, errors.New("Could not find repositoties under provided path")
	}

	color.Green("Found %v repositories.", len(repositories))

	return repositories, nil
}

func GetRepositoryFromPath(path string) (models.Repository, error) {
	isGitRepository := checkIfGitDirectoryExists(path)
	if !isGitRepository {
		return models.Repository{}, errors.New("not a git repository")
	}

	remote := getOrigin(path)
	return models.Repository{
		Origin: remote,
		Path:   path,
	}, nil
}
