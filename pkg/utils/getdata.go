package utils

import (
	"strconv"

	"github.com/gocolly/colly"
)

type Host struct {
	name   string
	status bool
	uptime string
	users  uint8
	load   float32
}

func toHost(row *colly.HTMLElement) (Host, error) {
	// get raw data
	var rawData [5]string
	row.ForEach("td", func(i int, col *colly.HTMLElement) {
		rawData[i] = col.Text
	})

	// convert to host
	name := rawData[0]
	status := rawData[1] == "up"
	uptime := rawData[2]
	var users uint8
	var load float32
	if status {
		convUsers, err := strconv.ParseUint(rawData[3], 10, 8)
		if err != nil {
			return Host{}, err
		}
		users = uint8(convUsers)

		convLoad, err := strconv.ParseFloat(rawData[4], 32)
		if err != nil {
			return Host{}, err
		}
		load = float32(convLoad)
	}
	return Host{name, status, uptime, users, load}, nil
}

func Scrape() ([]Host, error) {
	c := colly.NewCollector()
	var rows []Host
	var rowCount uint8
	c.OnHTML("tr", func(row *colly.HTMLElement) {
		rowCount++
		// skip first 3 <tr> elements
		if rowCount > 3 {
			host, _ := toHost(row)
			rows = append(rows, host)
		}
	})

	err := c.Visit("https://apps.cs.utexas.edu/unixlabstatus/")
	if err != nil {
		return rows, err
	}
	return rows, nil
}
