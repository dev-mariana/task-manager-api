package services

import (
	"context"

	"github.com/dev-mariana-task-manager-api/internal/entities"
	"github.com/dev-mariana-task-manager-api/internal/repositories"
)

type GetAllTasksService struct {
	repository repositories.ITaskRepository
}

func NewGetAllTasksService(repository repositories.ITaskRepository) *GetAllTasksService {
	return &GetAllTasksService{
		repository: repository,
	}
}

func (s *GetAllTasksService) GetAll(ctx context.Context) ([]*entities.Task, error) {
	tasks, err := s.repository.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return tasks, nil
}
