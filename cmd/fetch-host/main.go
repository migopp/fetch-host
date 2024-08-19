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
		log.Fatal("Error getting config.json path:", err)
	}

	// check if config'd
	exists, err := setup.ConfigExists(configPath)
	if err != nil {
		log.Fatal("Error verifying config file:", err)
	}

	// create config
	if !exists {
		err := setup.CreateConfig(configPath)
		if err != nil {
			log.Fatal("Error creating config file:", err)
		}
	}

	// fetch config data
	configChan := make(chan setup.Config)
	errChan := make(chan error)
	go utils.GetConfigData(configPath, configChan, errChan)

	// scrape
	hosts, err := utils.Scrape()
	if err != nil {
		log.Fatal("Error scraping host stats:", err)
	}

	// find best
	best, err := utils.GetBest(hosts)
	if err != nil {
		log.Fatal("Error sorting hosts:", err)
	}

	select {
	case config := <-configChan:
		// format + output cmd
		cmd := fmt.Sprintf(config.SshTemplate, config.UtcsUsername, best)
		err = clipboard.Init()
		if err != nil {
			log.Fatal("Error creating clipboard:", err)
		}
		clipboard.Write(clipboard.FmtText, []byte(cmd))
	case err = <-errChan:
		log.Fatal("Error fetching config data:", err)
	}

	return
}
