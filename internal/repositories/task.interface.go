package repositories

import (
	"context"

	"github.com/dev-mariana-task-manager-api/internal/entities"
)

type ITaskRepository interface {
	Create(ctx context.Context, task *entities.Task) (*entities.Task, error)
}
