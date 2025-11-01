package routes

import (
	"github.com/dev-mariana-task-manager-api/internal/handlers"
	"github.com/dev-mariana-task-manager-api/internal/repositories"
	"github.com/dev-mariana-task-manager-api/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupTaskRoutes(router *gin.Engine, db *gorm.DB) {
	repo := repositories.NewTaskRepository(db)
	service := services.NewCreateTaskService(repo)
	handler := handlers.NewCreateTaskHandler(service)

	router.POST("/tasks", handler.CreateTask)
}
