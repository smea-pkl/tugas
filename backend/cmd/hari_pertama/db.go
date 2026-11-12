package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("pkl.db"), &gorm.Config{})
	if err != nil {
		return db, err
	}

	return db, nil
}
