package repman

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
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

func FindRepositories(path string) {
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		if !info.IsDir() {
			// not a directory, skiping
			return nil
		}

		fmt.Printf("Searching in %s", path)
		found := checkIfGitDirectoryExists(path)
		if found {
			fmt.Printf("Found git repository in %s", path)
		}

		return nil
	})
}
