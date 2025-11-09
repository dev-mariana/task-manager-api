package services

import (
	"context"

	"github.com/dev-mariana-task-manager-api/internal/entities"
	"github.com/dev-mariana-task-manager-api/internal/repositories"
)

type GetTaskService struct {
	repository repositories.ITaskRepository
}

func NewGetTaskService(repository repositories.ITaskRepository) *GetTaskService {
	return &GetTaskService{
		repository: repository,
	}
}

func (s *GetTaskService) GetByID(ctx context.Context, id string) (*entities.Task, error) {
	task, err := s.repository.GetByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return task, nil
}
