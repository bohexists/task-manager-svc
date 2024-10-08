package app

import (
	"context"
	"github.com/bohexists/task-manager-svc/domain"
	"github.com/bohexists/task-manager-svc/ports/outbound"
)

// TaskService interface now works with domain.Task instead of proto.Task
type TaskService struct {
	TaskRepo outbound.TaskRepository
}

// NewTaskService creates new TaskService
func NewTaskService(repo outbound.TaskRepository) *TaskService {
	return &TaskService{
		TaskRepo: repo,
	}
}

// TaskService interface now works with domain.Task instead of proto.Task
func (s *TaskService) CreateTask(ctx context.Context, title, description string) (int64, error) {

	task, err := domain.NewTask(title, description)
	if err != nil {
		return 0, err
	}
	return s.TaskRepo.CreateTask(task)
}

// TaskService interface now works with domain.Task instead of proto.Task
func (s *TaskService) GetTask(ctx context.Context, id int64) (*domain.Task, error) {
	return s.TaskRepo.GetTask(id)
}

// TaskService interface now works with domain.Task instead of proto.Task
func (s *TaskService) UpdateTask(ctx context.Context, id int64, title, description string) error {
	task, err := s.TaskRepo.GetTask(id)
	if err != nil {
		return err
	}

	if err := task.Update(title, description); err != nil {
		return err
	}

	return s.TaskRepo.UpdateTask(task)
}

// TaskService interface now works with domain.Task instead of proto.Task
func (s *TaskService) DeleteTask(ctx context.Context, id int64) error {
	return s.TaskRepo.DeleteTask(id)
}

// TaskService interface now works with domain.Task instead of proto.Task
func (s *TaskService) ListTasks(ctx context.Context) ([]*domain.Task, error) {
	return s.TaskRepo.ListTasks()
}
