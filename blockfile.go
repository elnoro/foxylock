package main

import (
	"os"
	"strings"
)

type FileList struct {
	hosts []string
}

func newInMemory(filename string) (*FileList, error) {
	contents, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	hosts := strings.Split(string(contents), "\n")

	return &FileList{hosts}, nil
}

func (lh *FileList) addLockedHosts(hosts []string) []string {
	for _, expected := range lh.hosts {
		present := false

		for _, written := range hosts {
			if written == expected {
				present = true
				break
			}
		}

		if !present {
			hosts = append(hosts, expected)
		}
	}

	return hosts
}

func (lh *FileList) updateHostLists(hosts []string) []string {
	return lh.removeUnlockedHosts(lh.addLockedHosts(hosts))
}

func (lh *FileList) removeUnlockedHosts(hosts []string) []string {
	var lockedHosts []string
	for _, written := range hosts {
		present := false

		for _, locked := range lh.hosts {
			if written == locked {
				present = true
				break
			}
		}

		if present {
			lockedHosts = append(lockedHosts, written)
		}
	}

	return lockedHosts
}
