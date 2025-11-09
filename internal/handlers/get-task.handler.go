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

func (h *GetTaskHandler) GetTask(c *gin.Context) {
	id := c.Param("id")
	task, err := h.service.GetByID(c.Request.Context(), id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}
