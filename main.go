package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const osHosts = "/etc/hosts"

func main() {

	hostsFile, err := newHostsFile(osHosts)
	if err != nil {
		log.Fatal(err)
	}
	lockedHosts, err := newInMemory(getLockFile())
	if err != nil {
		log.Fatal(err)
	}

	subCmd := getSubCommand(os.Args)
	switch subCmd {
	case "lock":
		hostsFile.update(lockedHosts)
	case "unlock":
		hostsFile.clear()
	default:
		log.Fatal("Unknown command", subCmd)

	}

	err = backUp(osHosts, getConfigFolder()+"/backup_etc_hosts")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("🦊 About to write a new " + osHosts + " file")
	fmt.Println(hostsFile.toString())
	f, err := os.Create(osHosts)
	if err != nil {
		log.Fatal(err)
	}

	f.WriteString(hostsFile.toString())
}

func getSubCommand(args []string) (subCmd string) {
	if len(args) < 2 {
		subCmd = "lock"
	} else {
		subCmd = args[1]
	}

	return
}

func getLockFile() string {
	return getConfigFolder() + "/foxylock.txt"
}

func getConfigFolder() string {
	home, err := os.UserHomeDir()
	if err == nil {
		return home + "/.config/foxylock"
	}

	return "."
}

func backUp(orig string, back string) error {
	input, err := ioutil.ReadFile(orig)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(back, input, 0644)
}
