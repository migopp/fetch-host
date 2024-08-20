package utils

import (
	"errors"
	"sort"
)

func GetBest(hosts []Host) (string, error) {
	// filter down hosts
	upHosts := make([]Host, len(hosts))
	for _, host := range hosts {
		if host.status {
			upHosts = append(upHosts, host)
		}
	}
	if len(upHosts) == 0 {
		return "", errors.New("No running hosts")
	}

	// sort by load, user count
	sort.Slice(upHosts, func(i, j int) bool {
		if upHosts[i].load == upHosts[j].load {
			return upHosts[i].users < upHosts[j].users
		}
		return upHosts[i].load < upHosts[j].load
	})

	// first fit :3
	return upHosts[0].name, nil
}
