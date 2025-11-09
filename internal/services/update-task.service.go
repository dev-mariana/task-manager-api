package services

import (
	"context"
	"errors"

	"github.com/dev-mariana-task-manager-api/internal/entities"
	"github.com/dev-mariana-task-manager-api/internal/repositories"
)

type UpdateTaskService struct {
	repository repositories.ITaskRepository
}

func NewUpdateTaskService(repository repositories.ITaskRepository) *UpdateTaskService {
	return &UpdateTaskService{
		repository: repository,
	}
}

func (s *UpdateTaskService) Update(ctx context.Context, id string, dto *entities.UpdateTaskDTO) (*entities.Task, error) {
	updates := make(map[string]interface{})

	if dto.Title != nil {
		updates["title"] = *dto.Title
	}

	if dto.Description != nil {
		updates["description"] = *dto.Description
	}

	if dto.Status != nil {
		updates["status"] = *dto.Status
	}

	if len(updates) == 0 {
		return nil, errors.New("no fields to update")
	}

	updatedTask, err := s.repository.Update(ctx, id, updates)

	if err != nil {
		return nil, err
	}

	return updatedTask, nil
}
