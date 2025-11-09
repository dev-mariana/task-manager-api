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

// CreateTask godoc
// @Summary Create a new task
// @Description Create a new task with title, description, and status
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body entities.Task true "Task object"
// @Success 201 {object} entities.Task "Created task"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /tasks [post]
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
