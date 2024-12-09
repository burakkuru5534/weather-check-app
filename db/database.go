package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"weather-check-app/models"
)

var DB *gorm.DB

func InitDB(database string) {
	var err error
	DB, err = gorm.Open(sqlite.Open(database), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB.AutoMigrate(&models.WeatherQuery{})
}
