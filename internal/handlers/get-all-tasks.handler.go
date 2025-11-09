package handlers

import (
	"net/http"

	"github.com/dev-mariana-task-manager-api/internal/services"
	"github.com/gin-gonic/gin"
)

type GetAllTasksHandler struct {
	service *services.GetAllTasksService
}

func NewGetAllTasksHandler(service *services.GetAllTasksService) *GetAllTasksHandler {
	return &GetAllTasksHandler{service: service}
}

// GetAllTasks godoc
// @Summary Get all tasks
// @Description Retrieve a list of all tasks
// @Tags tasks
// @Produce json
// @Success 200 {array} entities.Task "List of tasks"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /tasks [get]
func (h *GetAllTasksHandler) GetAllTasks(c *gin.Context) {
	tasks, err := h.service.GetAll(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}
