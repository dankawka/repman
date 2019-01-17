package settingsmanager

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path"

	"github.com/dankawka/repman/internal/pkg/repofinder"
	"github.com/fatih/color"
)

func getAppStoreDirectory() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	appPath := path.Join(usr.HomeDir, ".repman")

	_, err = os.Stat(appPath)
	if err != nil {
		err = os.MkdirAll(appPath, 0755)
		if err != nil {
			color.Red("Could not create application folder %v", appPath)
		}
	}
	return appPath
}

func SaveRepositories(repositories []repofinder.Repository) {
	data, _ := json.Marshal(repositories)

	homePath := getAppStoreDirectory()
	filePath := path.Join(homePath, "repos.json")
	d1 := []byte(data)
	err := ioutil.WriteFile(filePath, d1, 0644)

	if err != nil {
		color.Red("Could not save file under %v", filePath)
	}
}
