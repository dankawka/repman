package settingsmanager

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/user"
	"path"

	"github.com/dankawka/repman/internal/pkg/models"
	"github.com/fatih/color"
)

func getAppStoreDirectory() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	appPath := path.Join(usr.HomeDir, ".repman")

	_, err = os.Stat(appPath)
	if err != nil {
		err = os.MkdirAll(appPath, 0755)
		if err != nil {
			color.Red("Could not create application folder %v", appPath)
			return "", err
		}
	}
	return appPath, nil
}

func getReposFilePath() (string, error) {
	homePath, err := getAppStoreDirectory()

	if err != nil {
		return "", err
	}

	filePath := path.Join(homePath, "repos.json")

	return filePath, nil
}

func SaveRepositories(repositories []models.Repository) error {
	repoFilePath, err := getReposFilePath()

	if err != nil {
		return err
	}

	data, _ := json.Marshal(repositories)

	d1 := []byte(data)
	err = ioutil.WriteFile(repoFilePath, d1, 0644)

	if err != nil {
		color.Red("Could not save file under %v", repoFilePath)
		return err
	}

	return nil
}

func GetListOfRepositories() ([]models.Repository, error) {
	repoFilePath, err := getReposFilePath()

	if err != nil {
		return nil, err
	}

	var result []models.Repository

	_, err = os.Stat(repoFilePath)
	if err != nil {
		return result, nil
	}

	jsonFile, err := os.Open(repoFilePath)
	defer jsonFile.Close()

	if err != nil {
		color.Red("Could not open file with list of repositories under %s", repoFilePath)
		return result, err
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal([]byte(byteValue), &result)
	return result, nil
}

func CheckIfAlreadySaved(repository models.Repository) bool {
	repositories, _ := GetListOfRepositories()

	exists := false

	for _, repo := range repositories {
		if repo.Origin == repository.Origin && repo.Path == repository.Path {
			exists = true
		}
	}

	return exists
}

func AppendRepository(repository models.Repository) error {
	alreadySavedRepositories, _ := GetListOfRepositories()

	extended := append(alreadySavedRepositories, repository)
	err := SaveRepositories(extended)

	if err != nil {
		return err
	}

	return nil
}
