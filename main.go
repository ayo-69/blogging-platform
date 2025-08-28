package main

import (
	"os"

	"github.com/ayo-69/blogging-platform/database"
	"github.com/ayo-69/blogging-platform/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	PORT, found := os.LookupEnv("PORT")
	if !found {
		PORT = "8080"
	}
	r := gin.Default()

	database.ConnectDB()

	routes.RegisterAuthRoutes(r)
	routes.RegisterPostRoutes(r)
	routes.RegisterProfileRoutes(r)
	routes.RegisterTagsRoutes(r)

	r.Run(":" + PORT)
}
