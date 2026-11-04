package main

import (
	"strings"
)

var blacklist []string = []string{
	"sutarman",
	"andika",
	"fitri",
}

func FilterUser(usr User) (bool, error) {
	username := strings.ToLower(usr.Username)
	gender := strings.ToLower(usr.Gender)

	for _, v := range blacklist {
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
