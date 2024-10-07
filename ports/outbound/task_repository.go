package outbound

import (
	"github.com/bohexists/task-manager-svc/api/proto"
)

type TaskRepository interface {
	CreateTask(task *proto.Task) (int64, error)
	GetTask(id int64) (*proto.Task, error)
	UpdateTask(task *proto.Task) error
	DeleteTask(id int64) error
	ListTasks() ([]*proto.Task, error)
}
