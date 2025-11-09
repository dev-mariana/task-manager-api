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

// DeleteTask godoc
// @Summary Delete a task
// @Description Delete a task by its ID
// @Tags tasks
// @Produce json
// @Param id path string true "Task ID"
// @Success 200 {object} map[string]string "Success message"
// @Failure 404 {object} map[string]string "Task not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /tasks/{id} [delete]
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
