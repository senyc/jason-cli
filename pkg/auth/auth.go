package auth

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/senyc/jason-cli/pkg/types"
)
func AddKeyToFS(key string) error {
	home, err :=os.UserHomeDir()
	if err != nil {
		return err
	}
	configFileName := fmt.Sprintf("%s/.config/jason/config.json", home)

	_, err = os.Stat(configFileName )
	if os.IsExist(err) {
		os.Remove(configFileName)
	} else {
		err = os.MkdirAll(fmt.Sprintf("%s/.config/jason", home), 0755)
		if err != nil {
			return err
		}
	}
	fileContents := types.ApiConfigFile{Key:key}
	j, err := json.Marshal(fileContents)
	if err != nil {
		return err
	}
	err = os.WriteFile(configFileName, j, 0644)
	return err
}

func GetKeyFromFS() (string, error){
	var key string
	var apiKeyFileContents types.ApiConfigFile
	home, err :=os.UserHomeDir()
	if err != nil {
		return key, err
	}

	file, err := os.ReadFile(fmt.Sprintf("%s/.config/jason/config.json", home))

	// This could include the fact that no file exists
	if err != nil {
		return key, err 
	}

	err = json.Unmarshal(file, &apiKeyFileContents)
	return apiKeyFileContents.Key, err
}
