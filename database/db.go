package database

import (
	"log"

	"github.com/ayo-69/blogging-platform/config"
	"github.com/ayo-69/blogging-platform/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(postgres.Open(config.DB_URL), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	log.Println("Connected to database")

	// Migrate the database schema
	DB.AutoMigrate(&models.Post{}, &models.User{}, &models.Tag{})
}
