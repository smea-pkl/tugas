package main

import (
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type Blacklist struct {
	Blacklists []string `yaml:"blacklist"`
}

func FilterUser(usr User) (bool, error) {
	var daftar Blacklist
	username := strings.ToLower(usr.Username)
	gender := strings.ToLower(usr.Gender)

	yamlConfig, err := os.ReadFile("config.yaml")
	if err != nil {
		return true, err
	}

	err = yaml.Unmarshal(yamlConfig, &daftar)

	for _, v := range daftar.Blacklists {
		if v == username {
			if username == "fitri" && usr.Gender != "p" {
				return true, nil
			}
			return false, nil
		}
	}

	switch gender {
	case "p":
	case "l":
	default:
		return false, nil
	}

	return true, nil
}
