package main

import (
	"log"
	"os"

	"github.com/dev-mariana-task-manager-api/internal/config"
	"github.com/dev-mariana-task-manager-api/internal/entities"
	"github.com/dev-mariana-task-manager-api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file.")
	}

	config.ConnectDB()
}

func main() {
	r := gin.Default()

	config.DB.AutoMigrate(&entities.Task{})

	routes.SetupTaskRoutes(r, config.DB)

	port := os.Getenv("SERVER_PORT")

	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
