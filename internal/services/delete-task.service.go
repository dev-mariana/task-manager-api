package services

import (
	"context"

	"github.com/dev-mariana-task-manager-api/internal/repositories"
)

type DeleteTaskService struct {
	repository repositories.ITaskRepository
}

func NewDeleteTaskService(repository repositories.ITaskRepository) *DeleteTaskService {
	return &DeleteTaskService{
		repository: repository,
	}
}

func (s *DeleteTaskService) Delete(ctx context.Context, id string) error {
	return s.repository.Delete(ctx, id)
}
