package main

import (
	"log"
	"os"

	"github.com/dev-mariana-task-manager-api/internal/config"
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

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server is running...",
		})
	})

	port := os.Getenv("SERVER_PORT")

	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
