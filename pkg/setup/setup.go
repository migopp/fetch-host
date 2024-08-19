package setup

import (
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	UtcsUsername string
	SshTemplate  string
}

func GetConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, ".config", "fetch-host", "config.json"), nil
}

func ConfigExists(configPath string) (bool, error) {
	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func CreateConfig(configPath string) error {
	// config template
	const configTemplate string = "{\n\t\"utcsUsername\": \"%s\",\n\t\"sshTemplate\": \"ssh %%s@%%s.cs.utexas.edu\"\n}"

	// create config directory
	parentFolder := filepath.Dir(configPath)
	err := os.MkdirAll(parentFolder, 0711)
	if err != nil {
		return err
	}

	// get config info
	fmt.Print("`fetch-host` setup\nUTCS Username: ")
	var utcsUsername string
	_, err = fmt.Scan(&utcsUsername)
	if err != nil {
		return err
	}
	filledConfig := fmt.Sprintf(configTemplate, utcsUsername)

	// write config to file
	err = os.WriteFile(configPath, []byte(filledConfig), 0644)
	if err != nil {
		return err
	}

	fmt.Println("Config created @", configPath)
	return nil
}
