package main

import (
	"fmt"
	"golang.design/x/clipboard"
	"log"

	"github.com/migopp/fetch-host/pkg/setup"
	"github.com/migopp/fetch-host/pkg/utils"
)

func main() {
	// get config path
	configPath, err := setup.GetConfigPath()
	if err != nil {
		log.Fatal("Error:", err)
	}

	// check if config'd
	exists, err := setup.ConfigExists(configPath)
	if err != nil {
		log.Fatal("Error:", err)
	}

	// create config
	if !exists {
		err := setup.CreateConfig(configPath)
		if err != nil {
			log.Fatal("Error:", err)
		}
	}

	// scrape
	hosts, err := utils.Scrape()
	if err != nil {
		log.Fatal("Error:", err)
	}

	// find best
	best, err := utils.GetBest(hosts)
	if err != nil {
		log.Fatal("Error:", err)
	}

	// format + output cmd
	config, err := utils.GetConfigData(configPath)
	if err != nil {
		log.Fatal("Error:", err)
	}
	cmd := fmt.Sprintf(config.SshTemplate, config.UtcsUsername, best)
	err = clipboard.Init()
	if err != nil {
		log.Fatal("Error:", err)
	}
	clipboard.Write(clipboard.FmtText, []byte(cmd))

	return
}
