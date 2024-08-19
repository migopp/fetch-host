package utils

import (
	"encoding/json"
	"os"

	"github.com/migopp/fetch-host/pkg/setup"
)

func GetConfigData(configPath string) (setup.Config, error) {
	// get file contents
	configFile, err := os.ReadFile(configPath)
	if err != nil {
		return setup.Config{}, err
	}

	// unmarshall data
	var config setup.Config
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		return setup.Config{}, err
	}

	return config, nil
}
