package handlers

import (
	"github.com/dev-mariana-task-manager-api/internal/services"
	"github.com/gin-gonic/gin"
)

type FindAllTasksHandler struct {
	service *services.GetAllTasksService
}

func NewFindAllTasksHandler(service *services.GetAllTasksService) *FindAllTasksHandler {
	return &FindAllTasksHandler{service: service}
}

func (h *FindAllTasksHandler) GetAllTasks(c *gin.Context) {
	tasks, err := h.service.GetAll(c.Request.Context())

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, tasks)
}
