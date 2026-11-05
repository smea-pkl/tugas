package main

import (
	"crypto/md5"
	"encoding/hex"
	"os"

	"gopkg.in/yaml.v3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	Database string `yaml:"database"`
}

func AddData(usr *User) (bool, error) {
	// var dbAnswer, fileAnswer string

	// fmt.Print("Mau disimpan ke DB (Y/n) : ")
	// fmt.Scan(&dbAnswer)
	// fmt.Print("Mau disimpan ke Answer (Y/n) : ")
	// fmt.Scan(&fileAnswer)

	// if strings.ToLower(dbAnswer) != "y" && strings.ToLower(fileAnswer) != "y" {
	// 	return false, nil
	// }

	// if dbAnswer == "Y" || dbAnswer == "y" {
	// 	dbfilename := "pkl.db"
	// 	db, err := gorm.Open(sqlite.Open(dbfilename), &gorm.Config{})

	// 	if err != nil {
	// 		return false, err
	// 	}

	// 	err = db.AutoMigrate(&User{})
	// 	if err != nil {
	// 		return false, err
	// 	}

	// 	err = db.Save(&usr).Error
	// 	if err != nil {
	// 		return false, err
	// 	}
	// }

	// if fileAnswer == "Y" || fileAnswer == "y" {
	// 	filename := "data.txt"

	// 	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	// 	if err != nil {
	// 		return false, err
	// 	}

	// 	defer file.Close()

	// 	file.WriteString(usr.UserToString())
	// }

	// return true, nil

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
		dbfilename := "pkl.db"
		db, err := gorm.Open(sqlite.Open(dbfilename), &gorm.Config{})

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
