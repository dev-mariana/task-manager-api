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

// UpdateTask godoc
// @Summary Update a task
// @Description Update an existing task by ID. All fields are optional.
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Param task body entities.UpdateTaskDTO true "Task update object"
// @Success 200 {object} entities.Task "Updated task"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /tasks/{id} [patch]
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
