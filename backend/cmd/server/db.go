package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		return db, err
	}

	err = db.AutoMigrate(User{})
	if err != nil {
		return db, err
	}

	return db, nil
}
