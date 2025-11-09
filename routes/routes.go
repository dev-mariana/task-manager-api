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
	createTaskService := services.NewCreateTaskService(repo)
	createTaskHandler := handlers.NewCreateTaskHandler(createTaskService)

	getAllService := services.NewGetAllTasksService(repo)
	getAllHandler := handlers.NewGetAllTasksHandler(getAllService)

	getTaskService := services.NewGetTaskService(repo)
	getTaskHandler := handlers.NewGetTaskHandler(getTaskService)

	router.POST("/tasks", createTaskHandler.CreateTask)
	router.GET("/tasks", getAllHandler.GetAllTasks)
	router.GET("/tasks/:id", getTaskHandler.GetTask)
}
