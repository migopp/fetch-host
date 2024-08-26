package utils

import (
	"encoding/json"
	"os"

	"github.com/migopp/fetch-host/pkg/setup"
)

func GetConfigData(configPath string, configChan chan<- setup.Config, errChan chan<- error) {
	// get file contents
	configFile, err := os.ReadFile(configPath)
	if err != nil {
		errChan <- err
		return
	}

	// unmarshall data
	var config setup.Config
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		errChan <- err
		return
	}

	configChan <- config
}
