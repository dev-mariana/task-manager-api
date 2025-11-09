package repositories

import (
	"context"

	"github.com/dev-mariana-task-manager-api/internal/entities"
)

type ITaskRepository interface {
	Create(ctx context.Context, task *entities.Task) (*entities.Task, error)
	GetAll(ctx context.Context) ([]*entities.Task, error)
	GetByID(ctx context.Context, id string) (*entities.Task, error)
	Update(ctx context.Context, id string, updates map[string]interface{}) (*entities.Task, error)
	Delete(ctx context.Context, id string) error
}
