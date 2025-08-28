package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT   string
	DB_URL string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	PORT = os.Getenv("PORT")
	DB_URL = os.Getenv("DB_URL")
}
