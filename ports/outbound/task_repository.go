package outbound

import (
	"github.com/bohexists/task-manager-svc/domain"
)

// TaskRepository interface now works with domain.Task instead of proto.Task
type TaskRepository interface {
	CreateTask(task *domain.Task) (int64, error)
	GetTask(id int64) (*domain.Task, error)
	UpdateTask(task *domain.Task) error
	DeleteTask(id int64) error
	ListTasks() ([]*domain.Task, error)
}
