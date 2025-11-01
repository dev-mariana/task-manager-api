package repositories

import "github.com/dev-mariana-task-manager-api/internal/entities"

type ITaskRepository interface {
	Create(task entities.Task) (entities.Task, error)
}
