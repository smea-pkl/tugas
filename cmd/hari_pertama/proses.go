package main

import (
	"crypto/md5"
	"encoding/hex"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database string `yaml:"database"`
}

func InsertData(usr *User) (bool, error) {
	yamlConfig, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	var config Config

	err = yaml.Unmarshal(yamlConfig, &config)
	if err != nil {
		panic(err)
	}

	passByte := md5.Sum([]byte(usr.Password))
	usr.Password = hex.EncodeToString(passByte[:])

	switch config.Database {
	case "sqlite":
		db, err := ConnectDB()

		if err != nil {
			return false, err
		}

		err = db.AutoMigrate(&User{})
		if err != nil {
			return false, err
		}

		err = db.Save(&usr).Error
		if err != nil {
			return false, err
		}
	case "file":
		filename := "data.txt"

		file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			return false, err
		}

		defer file.Close()

		file.WriteString(usr.UserToString())
	default:
		return false, nil
	}

	return true, nil
}
