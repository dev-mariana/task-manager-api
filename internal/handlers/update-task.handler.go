package handlers

import (
	"net/http"

	"github.com/dev-mariana-task-manager-api/internal/entities"
	"github.com/dev-mariana-task-manager-api/internal/services"
	"github.com/gin-gonic/gin"
)

type UpdateTaskHandler struct {
	service *services.UpdateTaskService
}

func NewUpdateTaskHandler(service *services.UpdateTaskService) *UpdateTaskHandler {
	return &UpdateTaskHandler{service: service}
}

func (h *UpdateTaskHandler) UpdateTask(c *gin.Context) {
	id := c.Param("id")

	var updateDTO entities.UpdateTaskDTO

	if err := c.ShouldBindJSON(&updateDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTask, err := h.service.Update(c.Request.Context(), id, &updateDTO)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedTask)
}
