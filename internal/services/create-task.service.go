package services

import (
	"context"

	"github.com/dev-mariana-task-manager-api/internal/entities"
	"github.com/dev-mariana-task-manager-api/internal/repositories"
)

type CreateTaskService struct {
	repository repositories.ITaskRepository
}

func NewCreateTaskService(repository repositories.ITaskRepository) *CreateTaskService {
	return &CreateTaskService{
		repository: repository,
	}
}

func (s *CreateTaskService) Create(ctx context.Context, task *entities.Task) (*entities.Task, error) {
	createdTask, err := s.repository.Create(ctx, task)

	if err != nil {
		return nil, err
	}

	return createdTask, nil
}
