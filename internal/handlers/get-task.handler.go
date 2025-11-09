package handlers

import (
	"net/http"

	"github.com/dev-mariana-task-manager-api/internal/services"
	"github.com/gin-gonic/gin"
)

type GetTaskHandler struct {
	service *services.GetTaskService
}

func NewGetTaskHandler(service *services.GetTaskService) *GetTaskHandler {
	return &GetTaskHandler{service: service}
}

// GetTask godoc
// @Summary Get a task by ID
// @Description Retrieve a specific task by its ID
// @Tags tasks
// @Produce json
// @Param id path string true "Task ID"
// @Success 200 {object} entities.Task "Task object"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /tasks/{id} [get]
func (h *GetTaskHandler) GetTask(c *gin.Context) {
	id := c.Param("id")
	task, err := h.service.GetByID(c.Request.Context(), id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}
