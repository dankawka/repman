package settingsmanager

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/user"
	"path"

	"github.com/dankawka/repman/internal/pkg/repofinder"
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

func SaveRepositories(repositories []repofinder.Repository) error {
	data, _ := json.Marshal(repositories)

	homePath, err := getAppStoreDirectory()

	if err != nil {
		return err
	}

	filePath := path.Join(homePath, "repos.json")
	d1 := []byte(data)
	err = ioutil.WriteFile(filePath, d1, 0644)

	if err != nil {
		color.Red("Could not save file under %v", filePath)
		return err
	}

	return nil
}

func GetListOfRepositories() ([]repofinder.Repository, error) {
	homePath, err := getAppStoreDirectory()
	var result []repofinder.Repository

	if err != nil {
		return result, err
	}

	filePath := path.Join(homePath, "repos.json")

	_, err = os.Stat(filePath)
	if err != nil {
		color.Red("File with list of repositories does not exists or is not accessible, use 'scan' option first.")
		return result, err
	}

	jsonFile, err := os.Open(filePath)
	defer jsonFile.Close()

	if err != nil {
		color.Red("Could not open file with list of repositories under %s", filePath)
		return result, err
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal([]byte(byteValue), &result)
	return result, nil
}
