package database

import (
	"log/slog"
	"pkl/backend/configs"
	"pkl/backend/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB(cfg configs.AppConfig) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(cfg.DatabaseFile), &gorm.Config{})
	if err != nil {
		return db, err
	}

	if cfg.AutoMigrate {
		slog.Info("Auto migrate Database !")
		err = db.AutoMigrate(
			&models.User{},
		)
		if err != nil {
			return db, err
		}
	}

	return db, nil
}
