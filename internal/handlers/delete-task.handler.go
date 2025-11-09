package handlers

import (
	"net/http"

	"github.com/dev-mariana-task-manager-api/internal/services"
	"github.com/gin-gonic/gin"
)

type DeleteTaskHandler struct {
	service *services.DeleteTaskService
}

func NewDeleteTaskHandler(service *services.DeleteTaskService) *DeleteTaskHandler {
	return &DeleteTaskHandler{service: service}
}

func (h *DeleteTaskHandler) DeleteTask(c *gin.Context) {
	id := c.Param("id")

	err := h.service.Delete(c.Request.Context(), id)

	if err != nil {
		if err.Error() == "task not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
