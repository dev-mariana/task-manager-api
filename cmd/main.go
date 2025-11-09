// @title           Task Manager API
// @version         1.0
// @description     A RESTful API for managing tasks
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

package main

import (
	"log"
	"os"

	docs "github.com/dev-mariana-task-manager-api/docs"

	"github.com/dev-mariana-task-manager-api/internal/config"
	"github.com/dev-mariana-task-manager-api/internal/entities"
	"github.com/dev-mariana-task-manager-api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	docs.SwaggerInfo.BasePath = "/"

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run(":" + port)
}
