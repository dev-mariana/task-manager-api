package handlers

import (
	"net/http"

	"github.com/dev-mariana-task-manager-api/internal/entities"
	"github.com/dev-mariana-task-manager-api/internal/services"
	"github.com/gin-gonic/gin"
)

type CreateTaskHandler struct {
	service *services.CreateTaskService
}

func NewCreateTaskHandler(service *services.CreateTaskService) *CreateTaskHandler {
	return &CreateTaskHandler{service: service}
}

func (h *CreateTaskHandler) CreateTask(c *gin.Context) {
	var task entities.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTask, err := h.service.Create(c.Request.Context(), &task)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdTask)
}
