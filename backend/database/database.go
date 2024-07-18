package database

import (
	"github.com/codescalersinternships/SecretNote-API-SPA-Mohamed-Riyad/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("myApp.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.User{}, &models.Note{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
