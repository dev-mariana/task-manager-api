package repositories

import (
	"context"

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

func (r *TaskRepository) Create(ctx context.Context, task *entities.Task) (*entities.Task, error) {
	if task.ID == "" {
		generate, _ := cuid2.Init(cuid2.WithLength(16))
		task.ID = generate()
	}

	if err := r.db.WithContext(ctx).Create(task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

func (r *TaskRepository) GetAll(ctx context.Context) ([]*entities.Task, error) {
	var tasks []*entities.Task

	if err := r.db.WithContext(ctx).Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *TaskRepository) GetByID(ctx context.Context, id string) (*entities.Task, error) {
	var task entities.Task

	if err := r.db.WithContext(ctx).First(&task, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &task, nil
}
