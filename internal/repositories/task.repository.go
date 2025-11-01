package repositories

import (
	"github.com/dev-mariana-task-manager-api/internal/entities"
	"github.com/nrednav/cuid2"
	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) Create(task entities.Task) (entities.Task, error) {
	if task.ID == "" {
		task.ID = cuid2.Generate()
	}

	if err := r.db.Create(&task).Create(task).Error; err != nil {
		return entities.Task{}, err
	}

	return task, nil
}
