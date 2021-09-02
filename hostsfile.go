package main

import (
	"os"
	"strings"
)

type Hostsfile struct {
	otherLines  []string
	lockedHosts []string
}

func newHostsFile(filename string) (*Hostsfile, error) {
	contents, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	return newHostsFileFromContents(string(contents))

}

func newHostsFileFromContents(contents string) (*Hostsfile, error) {
	hostLines := strings.Split(contents, "\n")

	var foxyLockHosts []string
	var otherLines []string
	for _, line := range hostLines {
		if strings.Contains(line, "# added by foxylock") {
			trimmed := strings.TrimSuffix(line, "# added by foxylock")
			parsedLine := strings.Fields(trimmed)
			if len(parsedLine) < 2 {
				otherLines = append(otherLines, line)
				continue
			}
			foxyLockHosts = append(foxyLockHosts, parsedLine[1])
		} else {
			otherLines = append(otherLines, line)
		}
	}

	return &Hostsfile{
		otherLines:  otherLines,
		lockedHosts: foxyLockHosts,
	}, nil
}

func (h *Hostsfile) update(lh LockedHosts) {
	h.lockedHosts = lh.updateHostLists(h.lockedHosts)
}

func (h *Hostsfile) clear() {
	h.lockedHosts = []string{}
}

func (h *Hostsfile) toString() string {
	var lockLines []string
	for _, flh := range h.lockedHosts {
		if flh == "" {
			continue
		}
		lockLines = append(lockLines, "127.0.0.1    "+flh+" # added by foxylock")
	}

	lines := append(h.otherLines, lockLines...)
	return strings.Join(lines, "\n")
}

type LockedHosts interface {
	updateHostLists([]string) []string
}
