package entities

import "time"

type Status string

const (
	InProgress Status = "In Progress"
	Pending    Status = "Pending"
	Completed  Status = "Completed"
)

type Task struct {
	ID          string    `gorm:"primaryKey;type:char(16)"`
	Title       string    `gorm:"type:varchar(100);not null" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	Status      Status    `gorm:"type:varchar(20);default:'Pending'"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
